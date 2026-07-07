package model

import "github.com/mangk/adminBox/config"

type SysRole struct {
	Model
	Name        string     `json:"name" gorm:"size:50;comment:角色名称"`
	Description string     `json:"description" gorm:"size:255;comment:角色简介"`
	UserList    []*SysUser `json:"user_list,omitempty" gorm:"many2many:sys_user_role"`
}

func (SysRole) TableName() string {
	return config.DBCfg()["default"].Prefix + "sys_role"
}
