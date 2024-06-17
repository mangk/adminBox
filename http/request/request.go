package request

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/admin/model"
)

const ContextPublicRequestKey = "__public_request__"
const ContextUserKey = "__user__"

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
func JWTUserId(ctx *gin.Context) int {
	return ctx.GetInt(ContextUserKey)
}

// 结合数据库获取用户最新信息
func JWTUserFetch(ctx *gin.Context) model.SysUser {
	u, _ := model.SysUser{}.Detail(ctx.MustGet(ContextUserKey).(int))
	return u
}
