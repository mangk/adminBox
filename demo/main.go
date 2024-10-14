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
	router.GET("component/t1", func(ctx *gin.Context) {
		ctx.String(200, "%s", components.T1)
	})
	router.GET("component/t2", func(ctx *gin.Context) {
		ctx.String(200, "%s", components.T2)
	})

	front.SetAdminBoxJsUserCodeSnippet(`
    Name: 'GOSKI',
    Host: '',
	Logo:'https://ch.goski.cn/1.ico',
	`, "")

	// front.RewriteIndex(func(ctx *gin.Context) {
	// 	ctx.JSON(200, 11111)
	// })

	s.ListenAndServer()
}
