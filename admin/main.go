package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/gAdmin/config"
	"github.com/mangk/gAdmin/http"
	"github.com/mangk/gAdmin/log"
)

func main() {
	s := http.New()
	h := http.HttpEngine()
	h.GET("/", func(ctx *gin.Context) {
		log.Infof("%s", config.Get("server.port"))
		ctx.JSON(418, 2222)
	})
	s.ListenAdnServer()
}
