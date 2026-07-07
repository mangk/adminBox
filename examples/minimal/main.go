package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/httpx"
	"github.com/mangk/adminBox/response"
)

func main() {
	// 注册自定义路由
	httpx.SetRouter(func(root *gin.Engine) {
		root.GET("/ping", func(ctx *gin.Context) {
			response.OkWithMsg(ctx, "pong")
		})

		root.GET("/hello", func(ctx *gin.Context) {
			response.OkWithData(ctx, "Hello AdminBox!")
		})

		root.POST("/echo", func(ctx *gin.Context) {
			var body map[string]interface{}
			ctx.ShouldBindJSON(&body)
			response.OkWithData(ctx, body)
		})
	})

	// 启动服务
	// 运行命令: go run main.go run
	httpx.Execute("minimal", "Minimal AdminBox Server")
}
