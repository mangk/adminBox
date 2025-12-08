package request

import (
	"github.com/gin-gonic/gin"
	abHttp "github.com/mangk/adminBox/http"
)

const ContextPublicRequestKey = "__public_request__"
const ContextLoginUserKey = "__login_user__"
const ContextRoleUserTypeKey = "__role_user_type__"
const ContextRoleUserIdKey = "__role_user_id__"

type PublicPageRequest struct {
	Query map[string]interface{} `json:"query,omitempty"`
	Sort  []string               `json:"sort,omitempty"`
	Id    string                 `json:"id,omitempty"`
	Ids   []any                  `json:"ids,omitempty"`
	abHttp.PageInfo
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

type PageQuery[T any] struct {
	Query    T               `json:"query"`
	PageInfo abHttp.PageInfo `json:"page_info"`
}

func PageRequest[T any](ctx *gin.Context, query T) (PageQuery[T], error) {
	req := PageQuery[T]{Query: query}
	err := ctx.ShouldBindJSON(&req)

	if req.PageInfo.Page <= 0 {
		req.PageInfo.Page = 1
	}
	if req.PageInfo.PageSize <= 0 {
		req.PageInfo.PageSize = 20
	}
	
	return req, err
}
