package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/http/request"
	"github.com/mangk/adminX/http/response"
	"io"
	"strings"
)

func PublicRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := request.PublicPageRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			if !errors.Is(err, io.EOF) {
				response.FailWithDetail(ctx, "请求出错", err.Error())
				ctx.Abort()
				return
			}
		}
		if req.Page == 0 {
			req.Page = 1
		}
		if req.PageSize == 0 {
			req.PageSize = 20
		}
		if req.PageSize > 500 {
			req.PageSize = 500
		}
		if req.Id != "" {
			for _, v := range strings.Split(req.Id, ",") {
				if strings.Contains(v, ",") {
					response.FailWithMsg(ctx, "读取数据 ID 错误")
				}
				if v != "" {
					req.Ids = append(req.Ids, v)
				}
			}
		}
		ctx.Set(request.ContextPublicRequestKey, req)
		ctx.Next()
	}

}
