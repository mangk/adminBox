package main

import (
	"demo/handler/components"
	"github.com/mangk/adminBox/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mangk/adminBox/front"
)

func main() {
	s := http.New()
	router := http.HttpEngine()
	router.GET("test", func(ctx *gin.Context) {
		ctx.String(200, "%s", components.Test)
	})

	s.ListenAndServer()
}
