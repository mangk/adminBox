package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mangk/adminX/cache"
	"github.com/mangk/adminX/db"
	"github.com/mangk/adminX/log"
	"github.com/mangk/adminX/utils"
	"gorm.io/gorm"
)

type SysUser struct {
	Model
	UUID           uuid.UUID        `json:"uuid" gorm:"index;size:60;comment:用户UUID"`
	Username       string           `json:"username" gorm:"index;size:30;comment:用户登录名"`
	Phone          string           `json:"phone"  gorm:"size:30;comment:用户手机号"`
	Email          string           `json:"email"  gorm:"size:60;comment:用户邮箱"`
	Salt           string           `json:"-" gorm:"size:16;comment:密码混淆"`
	Password       string           `json:"-"  gorm:"size:255;comment:用户登录密码"`
	NickName       string           `json:"nick_name" gorm:"size:30;default:系统用户;comment:用户昵称"`
	Avatar         string           `json:"avatar" gorm:"size:255;comment:用户头像"`
	Enable         bool             `json:"enable" gorm:"default:true;comment:用户是否有效"`
	UserConfig     UserConfig       `json:"user_config" gorm:"serializer:json;comment:用户配置文件"`
	ExtendConfig   string           `json:"extend_config" gorm:"comment:扩展配置，保存自定义使用的配置"`
	LastLoginAt    *LocalTime       `json:"last_login_at" gorm:"type:datetime;comment:最后登录时间"`
	DepartmentList []*SysDepartment `json:"department_list" gorm:"many2many:sys_user_department;"` // 用户部门
	DepartmentIds  []int            `json:"department_ids" gorm:"-"`                               // 用户部门 Id 集合
	RoleList       []*SysRole       `json:"role_list" gorm:"many2many:sys_user_role;"`             // 用户角色
	RoleIds        []int            `json:"role_ids" gorm:"-"`                                     // 用户角色 Id 集合
	JwtToken       string           `json:"jwt_token" gorm:"-"`
}

type UserConfig struct {
	SideMode    string `json:"side_mode"`
	BaseColor   string `json:"base_color"`
	ActiveColor string `json:"active_color"`
	HomePage    string `json:"home_page"`
}

func (s SysUser) cacheKey(id int) string {
	return fmt.Sprintf("%s:%d", s.TableName(), id)
}

func (s SysUser) TableName() string {
	return "sys_user"
}

func (s SysUser) Detail(id int) (user SysUser, err error) {
	data := cache.RedisHasOrQuery(s.cacheKey(id), func() string {
		if err = db.DB().Model(&user).Preload("DepartmentList").Preload("RoleList").Where("id = ?", id).First(&user).Error; err != nil {
			log.Error(err.Error())
			return ""
		}

		for _, v := range user.DepartmentList {
			user.DepartmentIds = append(user.DepartmentIds, v.ID)
		}

		for _, v := range user.RoleList {
			user.RoleIds = append(user.RoleIds, v.ID)
		}

		d, err := json.Marshal(user)
		if err != nil {
			log.Error(err.Error())
			return ""
		}

		return string(d)
	}, time.Hour*4)

	err = json.Unmarshal([]byte(data), &user)
	return
}

func (s SysUser) Update(data SysUser) error {
	return db.DB().Transaction(func(tx *gorm.DB) error {
		update := make(map[string]interface{})
		//update["ub"] = req.
		update["username"] = data.Username
		update["phone"] = data.Phone
		update["email"] = data.Email
		update["nick_name"] = data.NickName
		update["avatar"] = data.Avatar
		update["enable"] = data.Enable
		{
			if err := tx.Unscoped().Where("sys_user_id = ?", data.ID).Delete(&SysUserDepartment{}).Error; err != nil {
				return err
			}
			create := []SysUserDepartment{}
			for _, v := range data.DepartmentIds {
				create = append(create, SysUserDepartment{
					Model:           Model{Cb: data.Cb},
					SysUserId:       data.ID,
					SysDepartmentId: v,
				})
			}

			if len(create) > 0 {
				if err := tx.Create(&create).Error; err != nil {
					return err
				}
			}
		}

		{
			if err := tx.Unscoped().Where("sys_user_id = ?", data.ID).Delete(&SysUserRole{}).Error; err != nil {
				return err
			}
			create := []SysUserRole{}
			for _, v := range data.RoleIds {
				create = append(create, SysUserRole{
					Model:     Model{Cb: data.Cb},
					SysUserId: data.ID,
					SysRoleId: v,
				})
			}
			if len(create) > 0 {
				if err := tx.Create(&create).Error; err != nil {
					return err
				}
			}
		}

		if err := tx.Model(&data).Where("id = ?", data.ID).Updates(update).Error; err != nil {
			return err
		}

		cache.RedisDel(s.cacheKey(data.ID))
		return nil
	})
}

func (s SysUser) Login(username, password string) (user SysUser, err error) {
	if err = db.DB().Where("username = ?", username).Where("enable = ?", true).Preload("DepartmentList").Preload("RoleList").First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	password = user.Salt + password + user.Salt + user.Salt + user.Salt // 密码混淆规则

	if ok := utils.BcryptCheck(password, user.Password); !ok {
		err = errors.New("用户名或密码错误")
		return
	}

	if !user.Enable {
		err = errors.New("用户被禁止登录")
		return
	}

	user.LastLoginAt = LocalTime{}.Now()
	if err = db.DB().Select("last_login_at").Updates(user).Error; err != nil {
		return
	}

	return
}
