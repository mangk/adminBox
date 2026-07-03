package httpx

import "github.com/gin-gonic/gin"

type PageInfo struct {
	PageSize int64 `json:"page_size,omitempty"`
	Page     int64 `json:"page,omitempty"`
	Total    int64 `json:"total,omitempty"`
}

type Req[T any] struct {
	Data     T              `json:"data,omitempty"`
	Query    map[string]any `json:"query,omitempty"`
	PageInfo PageInfo       `json:"page_info,omitempty"`
}

func RequestPublic[T any](ctx *gin.Context) (req Req[T], err error) {
	req = Req[T]{}
	err = ctx.ShouldBindJSON(&req)
	return
}

func Request[T any](ctx *gin.Context) (req T, err error) {
	err = ctx.ShouldBindJSON(&req)
	return
}
