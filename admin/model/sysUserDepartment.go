package model

import "github.com/mangk/adminBox/config"

type SysUserDepartment struct {
	SysUserId       int `json:"user_id" gorm:"type:int(11)"`
	SysDepartmentId int `json:"department_id" gorm:"type:int(11)"`
}

func (SysUserDepartment) TableName() string {
	return config.DBCfg()["default"].Prefix + "sys_user_department"
}
