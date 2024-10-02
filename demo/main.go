package main

import (
	"demo/handler/components"

	"github.com/mangk/adminBox/front"
	"github.com/mangk/adminBox/http"

	_ "github.com/mangk/adminBox/front"

	"github.com/gin-gonic/gin"
)

func main() {
	s := http.New()
	router := http.HttpEngine()
	router.GET("test", func(ctx *gin.Context) {
		ctx.String(200, "%s", components.Test)
	})

	front.SetAdminBoxJsUserCodeSnippet(`
	Name: 'GOSKI',
	`, "")

	// front.RewriteIndex(func(ctx *gin.Context) {
	// 	ctx.JSON(200, 11111)
	// })

	s.ListenAndServer()
}
