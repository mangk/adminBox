package response

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/http/request"
	"net/http"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	T    int64       `json:"t"`
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
	ctx.JSON(http.StatusOK, Response{
		Code: SuccessStatus,
		Msg:  MsgOK,
		T:    time.Now().Unix(),
	})
}

func OkWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code: SuccessStatus,
		Msg:  msg,
		T:    time.Now().Unix(),
	})
}

func OkWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: SuccessStatus,
		Msg:  MsgOK,
		Data: data,
		T:    time.Now().Unix(),
	})
}

func OkWithDetail(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: SuccessStatus,
		Msg:  msg,
		Data: data,
		T:    time.Now().Unix(),
	})
}

func OkWithPageData(ctx *gin.Context, count int64, data interface{}) {
	req := request.PublicRequest(ctx)
	ctx.JSON(http.StatusOK, Response{
		Code: SuccessStatus,
		Msg:  MsgOK,
		Data: PageData{
			List:     data,
			PageSize: int64(req.PageSize),
			Page:     int64(req.Page),
			Total:    count,
		},
		T: time.Now().Unix(),
	})
}

func Fail(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: FailStatus,
		Msg:  MsgFail,
		T:    time.Now().Unix(),
	})
}

func FailWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Response{
		Code: FailStatus,
		Msg:  msg,
		T:    time.Now().Unix(),
	})
}

func FailWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: FailStatus,
		Msg:  MsgFail,
		Data: data,
		T:    time.Now().Unix(),
	})
}

func FailWithDetail(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: FailStatus,
		Msg:  msg,
		Data: data,
		T:    time.Now().Unix(),
	})
}

func FailWithCode(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, Response{
		Code: code,
		Msg:  msg,
		T:    time.Now().Unix(),
	})
}
