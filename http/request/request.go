package request

import (
	"github.com/gin-gonic/gin"
)

const ContextPublicRequestKey = "__public_request__"
const ContextLoginUserKey = "__login_user__"
const ContextRoleUserTypeKey = "__role_user_type__"
const ContextRoleUserIdKey = "__role_user_id__"

type PublicPageRequest struct {
	Query    map[string]interface{} `json:"query,omitempty"`
	Sort     string                 `json:"sort,omitempty"`
	Page     int64                  `json:"page,omitempty"`
	PageSize int64                  `json:"page_size,omitempty"`
	Id       string                 `json:"id,omitempty"`
	Ids      []any                  `json:"ids,omitempty"`
}

func PublicRequest(ctx *gin.Context) PublicPageRequest {
	return ctx.MustGet(ContextPublicRequestKey).(PublicPageRequest)
}

func PublicRequestCrud(ctx *gin.Context) CRUDRequest {
	return ctx.MustGet(ContextPublicRequestKey).(CRUDRequest)
}

// 获取 JWT 中存储的用户信息
func JWTLoginUserId(ctx *gin.Context) int {
	return ctx.GetInt(ContextLoginUserKey)
}

// Deprecated: 身份与ID可能被客户端篡改，不安全
// 身份用户ID
func RoleUserId(ctx *gin.Context) int {
	return ctx.GetInt(ContextRoleUserIdKey)
}

// Deprecated: 身份与ID可能被客户端篡改，不安全
// 身份用类型
func RoleUserType(ctx *gin.Context) string {
	return ctx.GetString(ContextRoleUserTypeKey)
}
