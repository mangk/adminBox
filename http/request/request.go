package request

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/admin/model"
)

const ContextPublicRequestKey = "__public_request__"
const ContextLoginUserKey = "__login_user__"
const ContextRoleUserTypeKey = "__role_user_type__"
const ContextRoleUserIdKey = "__role_user_id__"

type PublicPageRequest struct {
	Query    map[string]interface{} `json:"query,omitempty"`
	Sort     string                 `json:"sort,omitempty"`
	Page     int                    `json:"page,omitempty"`
	PageSize int                    `json:"page_size,omitempty"`
	Id       string                 `json:"id,omitempty"`
	Ids      []string               `json:"ids,omitempty"`
}

func PublicRequest(ctx *gin.Context) PublicPageRequest {
	return ctx.MustGet(ContextPublicRequestKey).(PublicPageRequest)
}

// 获取 JWT 中存储的用户信息
func JWTLoginUserId(ctx *gin.Context) int {
	return ctx.GetInt(ContextLoginUserKey)
}

// 结合数据库获取用户最新信息
func JWTLoginUserFetch(ctx *gin.Context) model.SysUser {
	u, _ := model.SysUser{}.Detail(ctx.MustGet(ContextLoginUserKey).(int))
	return u
}

// 身份用户ID
func RoleUserId(ctx *gin.Context) int {
	return ctx.GetInt(ContextRoleUserIdKey)
}

// 身份用类型
func RoleUserType(ctx *gin.Context) string {
	return ctx.GetString(ContextRoleUserTypeKey)
}
