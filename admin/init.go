package admin

import (
	"net/http"
	"strings"

	"github.com/mangk/adminBox/admin/handler"
	"github.com/mangk/adminBox/admin/model"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/db"
	myHttp "github.com/mangk/adminBox/http"
	"github.com/mangk/adminBox/http/middleware"
	"github.com/mangk/adminBox/log"
	"github.com/mangk/adminBox/moduleRegister"
)

func init() {
	moduleRegister.ModuleAdd(admin{})
}

type admin struct{}

func (admin) InitModule() {
	// dbAutoMigrate()
	routerCreate()
}

func DBMigrate() {
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

	log.Infof("you can create with SQL: %s", `
INSERT INTO "sys_auth" ("id", "cb", "ub", "db", "ct", "ut", "dt", "table_id", "table_module", "type", "key", "set_value") VALUES (1, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-501', 1),(2, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-702', 1),(3, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-202', 1),(4, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-300', 1),(5, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-502', 1),(6, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-604', 1),(7, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-704', 1),(8, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-104', 1),(9, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-302', 1),(10, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-303', 1),(11, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-403', 1),(12, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-500', 1),(13, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-503', 1),(14, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-504', 1),(15, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-703', 1),(16, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-301', 1),(17, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-303', 1),(18, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-600', 1),(19, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-603', 1),(20, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-401', 1),(21, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-601', 1),(22, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-602', 1),(23, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-201', 1),(24, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-304', 1),(25, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-400', 1),(26, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-700', 1),(27, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-102', 1),(28, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-105', 1),(29, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-203', 1),(30, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-402', 1),(31, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-100', 1),(32, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-200', 1),(33, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-405', 1),(34, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-505', 1),(35, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-305', 1),(36, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-101', 1),(37, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-106', 1),(38, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-200', 1),(39, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-306', 1),(40, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-103', 1),(41, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-204', 1),(42, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-304', 1),(43, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-301', 1),(44, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-302', 1),(45, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-404', 1),(46, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-701', 1),(47, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'menu', '-300', 1),(48, NULL, NULL, NULL, NULL, NULL, NULL, 1, 'sys_role', 'api', '-100', 1);
INSERT INTO "sys_role" ("id", "cb", "ub", "db", "ct", "ut", "dt", "name", "description") VALUES (1, NULL, NULL, NULL, NULL, NULL, NULL, 'superAdmin', '');
INSERT INTO "sys_user" ("id", "cb", "ub", "db", "ct", "ut", "dt", "uuid", "username", "phone", "email", "salt", "password", "nick_name", "avatar", "enable", "user_config", "extend_config", "last_login_at") VALUES (1, NULL, NULL, NULL, NULL, NULL, NULL, '3a55a5c1-05e4-40f5-b068-d3a0c53b6595', 'super', '12345678910', 'adminX@adminX.go', '', '$2a$10$opw6AiSwuDY8fJ8cl6QxleZsoBoQmZbfu2suZGVM0TKLwJarUa1UG', 'super', '', 1, '{\"side_mode\":\"\",\"base_color\":\"\",\"active_color\":\"\"}', '', '2024-10-01 12:22:11');
INSERT INTO "sys_user_role" ("sys_user_id", "sys_role_id") VALUES (1, 1);
	`)
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
				fileUpload.StaticFS(cfg.PrefixPath, http.Dir(cfg.StorePath))
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
