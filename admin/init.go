package admin

import (
	"net/http"
	"strings"

	"github.com/mangk/adminX/admin/handler"
	"github.com/mangk/adminX/admin/model"
	"github.com/mangk/adminX/config"
	"github.com/mangk/adminX/db"
	myHttp "github.com/mangk/adminX/http"
	"github.com/mangk/adminX/http/middleware"
	"github.com/mangk/adminX/moduleRegister"
)

func init() {
	moduleRegister.ModuleAdd(admin{})
}

type admin struct{}

func (admin) InitModule() {
	// dbAutoMigrate()
	routerCreate()
}

func dbAutoMigrate() {
	db.DB().AutoMigrate(
		model.SysUser{},
		model.SysMenu{},
		model.SysApi{},
		model.SysRole{},
		model.SysDepartment{},
		model.SysAuth{},
		model.SysUserRole{},
		model.SysUserDepartment{},
		model.SysFileUpload{},
	)
}

func routerCreate() {
	router := myHttp.HttpEngine()

	router.Use(middleware.Cors()) // TODO 跨域限制基于配置

	backend := router.Group(strings.Trim(config.ServerCfg().BackendRouterPrefix, "/") + "/sys")
	backend.POST("login", handler.AuthLogin)
	backend.GET("logout", handler.AuthLogout)
	backend.GET("verificationCode", handler.AuthVerificationCode)

	backend.Use(middleware.JWTCheckByCasbin())

	auth := backend.Group("auth")
	{ // 用户授权信息
		auth.GET("userPermission", handler.AuthUserPermission)
		auth.GET("permissionAll", handler.AuthPermissionAll)
		auth.POST("permissionGetByIdAndModule", handler.AuthPermissionGetByIdAndModule)
		auth.PUT("permissionSave", handler.AuthPermissionSave)
	}

	{ // 文件管理
		fileUpload := backend.Group("fileUpload")
		fileUpload.POST("upload", handler.FileUpload)
		fileUpload.POST("page", middleware.PublicRequest(), handler.FileList)
		fileUpload.PUT("", handler.FileEdit)
		fileUpload.DELETE("", handler.FileDelete)
		fileUpload.GET("cfg", handler.FileGetUploadLimit)
		for _, cfg := range config.FileCfg() {
			if cfg.Driver == "local" {
				router.StaticFS(cfg.PrefixPath, http.Dir(cfg.StorePath))
			}
		}

	}

	setting := backend.Group("setting")
	{ // 菜单管理
		menu := setting.Group("menu")
		menu.POST("page", handler.Menu)
		menu.POST("getById", handler.MenuDetail)
		menu.POST("", handler.MenuCreate)
		menu.PUT("", handler.MenuEdit)
		menu.DELETE("", handler.MenuDelete)
	}

	{ // 接口管理
		api := setting.Group("api")
		api.POST("page", middleware.PublicRequest(), handler.Api)
		api.POST("getById", handler.ApiDetail)
		api.POST("", handler.ApiCreate)
		api.PUT("", handler.ApiEdit)
		api.DELETE("", handler.ApiDelete)
	}

	{ // 用户管理
		api := setting.Group("user")
		api.POST("page", middleware.PublicRequest(), handler.User)
		api.POST("getById", handler.UserDetail)
		api.POST("", handler.UserCreate)
		api.PUT("", handler.UserEdit)
		api.PUT("changePassword", handler.UserChangePassord)
		api.DELETE("", handler.UserDelete)
	}

	{ // 角色管理
		role := setting.Group("role")
		role.POST("page", middleware.PublicRequest(), handler.Role)
		role.POST("getById", handler.RoleDetail)
		role.POST("", handler.RoleCreate)
		role.PUT("", handler.RoleEdit)
		role.DELETE("", handler.RoleDelete)
		role.GET("all", handler.RoleAll)
	}

	{ // 部门管理
		department := setting.Group("department")
		department.POST("page", handler.Department)
		department.POST("getById", handler.DepartmentDetail)
		department.POST("", handler.DepartmentCreate)
		department.PUT("", handler.DepartmentEdit)
		department.DELETE("", handler.DepartmentDelete)
	}
}
