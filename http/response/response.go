package response

import (
	"errors"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	abHttp "github.com/mangk/adminBox/http"
	"github.com/mangk/adminBox/http/request"
)

var useResponseRecover atomic.Bool

// ResponseAbort is used ONLY to stop gin handler chain
// after response has been written.
// It MUST be recovered by middleware.ResponseRecover.
type ResponseAbort struct{}

func EnableResponseRecover() {
	useResponseRecover.Store(true)
}

func DisableResponseRecover() {
	useResponseRecover.Store(false)
}

func IsUseResponseRecover() bool {
	return useResponseRecover.Load()
}

type Response struct {
	Code           int         `json:"code"`
	Msg            string      `json:"msg"`
	MsgShowTimeout int64       `json:"msg_show_timeout,omitempty"`
	Data           interface{} `json:"data,omitempty"`
	T              int64       `json:"t"`
}

type PageResponse struct {
	List     any             `json:"list"`
	PageInfo abHttp.PageInfo `json:"page_info"`
}

type PageData struct {
	List     interface{} `json:"list"`
	PageSize int64       `json:"page_size"`
	Page     int64       `json:"page"`
	Total    int64       `json:"total"`
}

const (
	SuccessStatus = 0
	FailStatus    = -1

	MsgOK   = "ok"
	MsgFail = "fail"
)

// func ()  {

// }

func Ok(ctx *gin.Context) {
	jsonResponse(ctx, SuccessStatus, MsgOK, nil)
}

func OkWithMsg(ctx *gin.Context, msg string, showTime ...int64) {
	jsonResponse(ctx, SuccessStatus, msg, nil, showTime...)
}

func OkWithData(ctx *gin.Context, data interface{}) {
	jsonResponse(ctx, SuccessStatus, MsgOK, data)
}

func OkWithDetail(ctx *gin.Context, msg string, data interface{}, showTime ...int64) {
	jsonResponse(ctx, SuccessStatus, msg, data, showTime...)
}

func OkWithPageData(ctx *gin.Context, count int64, data interface{}) {
	req := request.PublicRequest(ctx)
	jsonResponse(ctx, SuccessStatus, MsgOK, PageData{
		List:     data,
		PageSize: int64(req.PageSize),
		Page:     int64(req.Page),
		Total:    count,
	})
}

func Fail(ctx *gin.Context) {
	jsonResponse(ctx, FailStatus, MsgFail, nil)
}

func FailWithMsg(ctx *gin.Context, msg string, showTime ...int64) {
	jsonResponse(ctx, FailStatus, msg, nil, showTime...)
}

func FailWithError(ctx *gin.Context, err error, showTime ...int64) {
	// TODO 提供错误提示转换map
	jsonResponse(ctx, FailStatus, err.Error(), nil, showTime...)
}

func FailWithData(ctx *gin.Context, data interface{}) {
	jsonResponse(ctx, FailStatus, MsgFail, data)
}

func FailWithDetail(ctx *gin.Context, msg string, data interface{}, showTime ...int64) {
	jsonResponse(ctx, FailStatus, msg, data, showTime...)
}

func FailWithCode(ctx *gin.Context, code int, msg string) {
	jsonResponse(ctx, code, msg, nil)
}

func FailWithCodeAndNeedReload(ctx *gin.Context, code int, msg string) {
	jsonResponse(ctx, code, msg, gin.H{"reload": true})
}


func FailAndAbort(ctx *gin.Context) {
	jsonResponseAndAbort(ctx, FailStatus, MsgFail, nil)
}

func FailWithMsgAndAbort(ctx *gin.Context, msg string, showTime ...int64) {
	jsonResponseAndAbort(ctx, FailStatus, msg, nil, showTime...)
}

func FailWithErrorAndAbort(ctx *gin.Context, err error, showTime ...int64) {
	// TODO 提供错误提示转换map
	jsonResponseAndAbort(ctx, FailStatus, err.Error(), nil, showTime...)
}

func FailWithDataAndAbort(ctx *gin.Context, data interface{}) {
	jsonResponseAndAbort(ctx, FailStatus, MsgFail, data)
}

func FailWithDetailAndAbort(ctx *gin.Context, msg string, data interface{}, showTime ...int64) {
	jsonResponseAndAbort(ctx, FailStatus, msg, data, showTime...)
}

func FailWithCodeAndAbort(ctx *gin.Context, code int, msg string) {
	jsonResponseAndAbort(ctx, code, msg, nil)
}

func FailWithCodeAndNeedReloadAndAbort(ctx *gin.Context, code int, msg string) {
	jsonResponseAndAbort(ctx, code, msg, gin.H{"reload": true})
}

func jsonResponse(ctx *gin.Context, code int, msg string, data interface{}, showTime ...int64) {
	var st int64 = 0
	if len(showTime) > 0 && showTime[0] > 0 {
		st = showTime[0]
	}

	if code != 0 {
		ctx.Error(errors.New(msg))
	}

	ctx.JSON(http.StatusOK, Response{
		Code:           code,
		Msg:            msg,
		MsgShowTimeout: st,
		Data:           data,
		T:              time.Now().Unix(),
	})
}
func jsonResponseAndAbort(ctx *gin.Context, code int, msg string, data interface{}, showTime ...int64) {
	var st int64 = 0
	if len(showTime) > 0 && showTime[0] > 0 {
		st = showTime[0]
	}

	if code != 0 {
		ctx.Error(errors.New(msg))
	}

	ctx.JSON(http.StatusOK, Response{
		Code:           code,
		Msg:            msg,
		MsgShowTimeout: st,
		Data:           data,
		T:              time.Now().Unix(),
	})
	if IsUseResponseRecover() {
		panic(ResponseAbort{})
	}
}

