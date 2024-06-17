package model

import (
	"github.com/mangk/adminX/config"
	"github.com/mangk/adminX/db"
)

type SysApi struct {
	Model
	MenuId      int    `json:"menu_id" gorm:"type:int(11);comment:菜单Id"`
	MenuName    string `json:"menu_name" gorm:"-"`
	Name        string `json:"name" gorm:"size:50;comment:api名称"`
	Description string `json:"description" gorm:"size:255;comment:api简介"`
	Path        string `json:"path" gorm:"size:255;comment:api路径"`
	Method      string `json:"method" gorm:"default:POST;comment:方法:创建POST(默认)|查看GET|更新PUT|删除DELETE"`

	Premiss int `json:"premiss,omitempty" gorm:"-"` // 权限列表展示使用
}

func (SysApi) TableName() string {
	return "sys_api"
}

func (s SysApi) All(loadSystem bool) []SysApi {
	list := []SysApi{}
	db.DB().Find(&list)

	if loadSystem {
		list = append(list, s.SysApi()...)
	}
	return list
}

func (s SysApi) SysApi() []SysApi {
	backendPrefix := config.ServerCfg().BackendRouterPrefix
	if backendPrefix != "" {
		backendPrefix = "/" + backendPrefix
	}

	return []SysApi{
		{Model: Model{ID: -100}, MenuId: -301, Name: "登录", Description: "", Path: backendPrefix + "/sys/login", Method: "POST"},
		{Model: Model{ID: -101}, MenuId: -301, Name: "退出登录", Description: "", Path: backendPrefix + "/sys/logout", Method: "GET"},
		{Model: Model{ID: -102}, MenuId: -301, Name: "验证码获取", Description: "", Path: backendPrefix + "/sys/verificationCode", Method: "GET"},

		{Model: Model{ID: -103}, MenuId: -301, Name: "用户授权列表", Description: "", Path: backendPrefix + "/sys/auth/userPermission", Method: "GET"},
		{Model: Model{ID: -104}, MenuId: -301, Name: "授权树", Description: "", Path: backendPrefix + "/sys/auth/permissionAll", Method: "GET"},
		{Model: Model{ID: -105}, MenuId: -301, Name: "授权详情", Description: "", Path: backendPrefix + "/sys/auth/permissionGetByIdAndModule", Method: "POST"},
		{Model: Model{ID: -106}, MenuId: -301, Name: "保存授权", Description: "", Path: backendPrefix + "/sys/auth/permissionSave", Method: "PUT"},

		{Model: Model{ID: -200}, MenuId: -302, Name: "菜单列表", Description: "", Path: backendPrefix + "/sys/setting/menu/page", Method: "POST"},
		{Model: Model{ID: -201}, MenuId: -302, Name: "菜单详情", Description: "", Path: backendPrefix + "/sys/setting/menu/getById", Method: "POST"},
		{Model: Model{ID: -202}, MenuId: -302, Name: "菜单创建", Description: "", Path: backendPrefix + "/sys/setting/menu", Method: "POST"},
		{Model: Model{ID: -203}, MenuId: -302, Name: "菜单更新", Description: "", Path: backendPrefix + "/sys/setting/menu", Method: "PUT"},
		{Model: Model{ID: -204}, MenuId: -302, Name: "菜单删除", Description: "", Path: backendPrefix + "/sys/setting/menu", Method: "DELETE"},

		{Model: Model{ID: -300}, MenuId: -303, Name: "API列表", Description: "", Path: backendPrefix + "/sys/setting/api/page", Method: "POST"},
		{Model: Model{ID: -301}, MenuId: -303, Name: "API详情", Description: "", Path: backendPrefix + "/sys/setting/api/getById", Method: "POST"},
		{Model: Model{ID: -302}, MenuId: -303, Name: "API创建", Description: "", Path: backendPrefix + "/sys/setting/api", Method: "POST"},
		{Model: Model{ID: -303}, MenuId: -303, Name: "API更新", Description: "", Path: backendPrefix + "/sys/setting/api", Method: "PUT"},
		{Model: Model{ID: -304}, MenuId: -303, Name: "API删除", Description: "", Path: backendPrefix + "/sys/setting/api", Method: "DELETE"},

		{Model: Model{ID: -400}, MenuId: -304, Name: "用户列表", Description: "", Path: backendPrefix + "/sys/setting/user/page", Method: "POST"},
		{Model: Model{ID: -401}, MenuId: -304, Name: "用户详情", Description: "", Path: backendPrefix + "/sys/setting/user/getById", Method: "POST"},
		{Model: Model{ID: -402}, MenuId: -304, Name: "用户创建", Description: "", Path: backendPrefix + "/sys/setting/user", Method: "POST"},
		{Model: Model{ID: -403}, MenuId: -304, Name: "用户更新", Description: "", Path: backendPrefix + "/sys/setting/user", Method: "PUT"},
		{Model: Model{ID: -404}, MenuId: -304, Name: "更改密码", Description: "", Path: backendPrefix + "/sys/setting/user/changePassword", Method: "PUT"},
		{Model: Model{ID: -405}, MenuId: -304, Name: "用户删除", Description: "", Path: backendPrefix + "/sys/setting/user", Method: "DELETE"},

		{Model: Model{ID: -500}, MenuId: -305, Name: "角色列表", Description: "", Path: backendPrefix + "/sys/setting/role/page", Method: "POST"},
		{Model: Model{ID: -501}, MenuId: -305, Name: "角色详情", Description: "", Path: backendPrefix + "/sys/setting/role/getById", Method: "POST"},
		{Model: Model{ID: -502}, MenuId: -305, Name: "角色创建", Description: "", Path: backendPrefix + "/sys/setting/role", Method: "POST"},
		{Model: Model{ID: -503}, MenuId: -305, Name: "角色更新", Description: "", Path: backendPrefix + "/sys/setting/role", Method: "PUT"},
		{Model: Model{ID: -504}, MenuId: -305, Name: "角色删除", Description: "", Path: backendPrefix + "/sys/setting/role", Method: "DELETE"},
		{Model: Model{ID: -505}, MenuId: -305, Name: "所有角色", Description: "", Path: backendPrefix + "/sys/setting/role/all", Method: "GET"},

		{Model: Model{ID: -600}, MenuId: -306, Name: "部门列表", Description: "", Path: backendPrefix + "/sys/setting/department/page", Method: "POST"},
		{Model: Model{ID: -601}, MenuId: -306, Name: "部门详情", Description: "", Path: backendPrefix + "/sys/setting/department/getById", Method: "POST"},
		{Model: Model{ID: -602}, MenuId: -306, Name: "部门创建", Description: "", Path: backendPrefix + "/sys/setting/department", Method: "POST"},
		{Model: Model{ID: -603}, MenuId: -306, Name: "部门更新", Description: "", Path: backendPrefix + "/sys/setting/department", Method: "PUT"},
		{Model: Model{ID: -604}, MenuId: -306, Name: "部门删除", Description: "", Path: backendPrefix + "/sys/setting/department", Method: "DELETE"},

		{Model: Model{ID: -700}, MenuId: -200, Name: "文件上传", Description: "", Path: backendPrefix + "/sys/fileUpload/upload", Method: "POST"},
		{Model: Model{ID: -701}, MenuId: -200, Name: "文件列表", Description: "", Path: backendPrefix + "/sys/fileUpload/page", Method: "POST"},
		{Model: Model{ID: -702}, MenuId: -200, Name: "编辑文件", Description: "", Path: backendPrefix + "/sys/fileUpload", Method: "PUT"},
		{Model: Model{ID: -703}, MenuId: -200, Name: "删除文件", Description: "", Path: backendPrefix + "/sys/fileUpload", Method: "DELETE"},
		{Model: Model{ID: -704}, MenuId: -200, Name: "文件上传配置", Description: "", Path: backendPrefix + "/sys/fileUpload/cfg", Method: "GET"},
	}
}
