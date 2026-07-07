package httpx

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ========== 响应状态码定义 ==========

const (
	ResponseStatus_OK   = 0
	ResponseStatus_Fail = -1
)

var ResponseStatusMap = map[int]string{
	ResponseStatus_OK:   "ok",
	ResponseStatus_Fail: "fail",
}

// ========== 配置常量 ==========

const (
	MaxDataSize     = 100 * 1024 * 1024 // 100MB
	MaxStringLength = 10 * 1024 * 1024  // 10MB
)

// ========== RespOption 类型定义 ==========

type RespOption func(*gin.Context, *Resp)

// ========== Resp 响应结构体 ==========

type Resp struct {
	httpCode int
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	Data     any       `json:"data,omitempty"`
	PageInfo *PageInfo `json:"page_info,omitempty"`
	T        int64     `json:"t"`

	error      error
	responseFn func(*gin.Context)
}

// SetResponseFunc 设置自定义响应函数，供子包选项使用
func (r *Resp) SetResponseFunc(fn func(*gin.Context)) {
	if r.responseFn != nil {
		panic("duplicate response function assignment")
	}
	r.responseFn = fn
}

// ========== 核心响应函数 ==========

func Response(ctx *gin.Context, opts ...RespOption) {
	r := &Resp{
		httpCode: http.StatusOK,
		Code:     ResponseStatus_OK,
		Msg:      ResponseStatusMap[ResponseStatus_OK],
		T:        time.Now().Unix(),
		error:    nil,
	}

	for _, opt := range opts {
		opt(ctx, r)
		if r.error != nil {
			r.Msg = r.error.Error()
			break
		}
	}

	if r.responseFn == nil {
		opt := AsInnerJson()
		opt(ctx, r)
	}

	r.responseFn(ctx)
}
