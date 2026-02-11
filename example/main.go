package main

import (
	_ "embed"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mangk/adminBox/pkg/admin"
	"github.com/mangk/adminBox/pkg/admin/front"
	"github.com/mangk/adminBox/pkg/httpServer"
	"github.com/mangk/adminBox/pkg/response"
)

//go:embed example.vue
var exampleTemplate string

func main() {
	httpServer.SetRouter(func(root *gin.Engine) {
		root.GET("example", front.TemplateBuild("example.vue", exampleTemplate))
		root.POST("api/example", func(ctx *gin.Context) {
			response.OkWithData(ctx, response.Response{
				Code: 0,
				Msg:  "ok",
				Data: "Hello AdminBox " + time.Now().Format("2006-01-02 15:04:05"),
			})
		})
	})

	httpServer.Execute("example", "Adminbox 的演示程序")
}
