package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/http/request"
)

type Response struct {
	Code           int         `json:"code"`
	Msg            string      `json:"msg"`
	MsgShowTimeout int64       `json:"msg_show_timeout,omitempty"`
	Data           interface{} `json:"data,omitempty"`
	T              int64       `json:"t"`
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

	MsgOK   = ""
	MsgFail = "fail"
)

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

func FailWithData(ctx *gin.Context, data interface{}) {
	jsonResponse(ctx, FailStatus, MsgFail, data)
}

func FailWithDetail(ctx *gin.Context, msg string, data interface{}, showTime ...int64) {
	jsonResponse(ctx, FailStatus, msg, data, showTime...)
}

func FailWithCode(ctx *gin.Context, code int, msg string) {
	jsonResponse(ctx, code, msg, nil)
}

func jsonResponse(ctx *gin.Context, code int, msg string, data interface{}, showTime ...int64) {
	var st int64 = 0
	if len(showTime) > 0 && showTime[0] > 0 {
		st = showTime[0]
	}

	ctx.JSON(http.StatusOK, Response{
		Code:           code,
		Msg:            msg,
		MsgShowTimeout: st,
		Data:           data,
		T:              time.Now().Unix(),
	})
}
