package httpx

import (
	"errors"
	"io"
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

// ========== 错误定义 ==========
type HttpxErr error

var (
	ErrInvalidHTTPCode       HttpxErr = errors.New("invalid HTTP status code")
	ErrInvalidContentType    HttpxErr = errors.New("invalid content type")
	ErrEmptyFilePath         HttpxErr = errors.New("file path cannot be empty")
	ErrInvalidTemplateName   HttpxErr = errors.New("template name cannot be empty")
	ErrNilReader             HttpxErr = errors.New("reader cannot be nil")
	ErrReaderAlreadyConsumed HttpxErr = errors.New("reader already consumed")
	ErrDataTooLarge          HttpxErr = errors.New("data too large")
	ErrCircularReference     HttpxErr = errors.New("circular reference detected in data")
	ErrInvalidUnicode        HttpxErr = errors.New("invalid unicode string")
	ErrHTMLInjection         HttpxErr = errors.New("potential HTML injection detected")
	ErrPathTraversal         HttpxErr = errors.New("path traversal detected")
)

// ========== 配置常量 ==========

const (
	MaxDataSize     = 100 * 1024 * 1024 // 100MB
	MaxStringLength = 10 * 1024 * 1024  // 10MB
)

// ========== RespOption 类型定义 ==========

type RespOption func(*Resp)

// ========== Resp 响应结构体 ==========

type Resp struct {
	httpCode int
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data,omitempty"`
	PageInfo *PageInfo   `json:"page_info,omitempty"`
	T        int64       `json:"t"`

	error error

	responseFn func(*gin.Context)

	reader        io.Reader
	contentType   string
	contentLength int
	extraHeaders  map[string]string
	cacheReader   bool
	cachedData    []byte
	filePath      string
	fileName      string
	htmlName      string
	htmlData      interface{}
	rawData       interface{}
}

// SetResponseFunc 设置自定义响应函数，供子包选项使用
func (r *Resp) SetResponseFunc(fn func(*gin.Context)) {
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
		opt(r)
		if r.error != nil {
			r.Msg = r.error.Error()
			break
		}
	}

	if r.responseFn != nil {
		r.responseFn(ctx)
		return
	}

	ctx.JSON(r.httpCode, *r)
}
