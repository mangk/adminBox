package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/pkg/httpServer"
	"github.com/mangk/adminBox/pkg/response"
)

func main() {
	// 注册自定义路由
	httpServer.SetRouter(func(root *gin.Engine) {
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
	httpServer.Execute("minimal", "Minimal AdminBox Server")
}
