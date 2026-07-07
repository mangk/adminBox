package model

import "github.com/mangk/adminBox/config"

type SysUserRole struct {
	SysUserId int `json:"user_id" gorm:"type:int(11)"`
	SysRoleId int `json:"role_id" gorm:"type:int(11)"`
}

func (SysUserRole) TableName() string {
	return config.DBCfg()["default"].Prefix + "sys_user_role"
}
