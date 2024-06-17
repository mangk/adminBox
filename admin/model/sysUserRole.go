package model

type SysUserRole struct {
	Model
	SysUserId int `json:"user_id" gorm:"type:int(11)"`
	SysRoleId int `json:"role_id" gorm:"type:int(11)"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
