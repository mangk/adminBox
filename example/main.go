package main

import (
	_ "embed"

	"github.com/gin-gonic/gin"
	_ "github.com/mangk/adminBox/pkg/admin/api"
	"github.com/mangk/adminBox/pkg/admin/front"
	"github.com/mangk/adminBox/pkg/httpServer"
	"github.com/mangk/adminBox/pkg/response"
)

//go:embed example.vue
var exampleTemplate string

func main() {
	httpServer.SetRouter(func(root *gin.Engine) {
		root.GET("example", front.TemplateBuild("example.vue", exampleTemplate))
		root.POST("example", func(ctx *gin.Context) {
			response.OkWithData(ctx, response.Response{
				Code: 0,
				Msg:  "ok",
				Data: "Hello AdminBox",
			})
		})
	})

	httpServer.Execute("example", "example")
}
