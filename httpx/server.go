package httpx

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
	"github.com/mangk/adminBox/middleware"
)

var _waitInitRoter = make([]func(root *gin.Engine), 0)
var _waitBrforeRun = make([]func(), 0)

func SetBeforeRun(f func()) {
	_waitBrforeRun = append(_waitBrforeRun, f)
}

func SetRouter(f func(root *gin.Engine)) {
	_waitInitRoter = append(_waitInitRoter, f)
}

func httpServer(cfgPath string) error {
	if cfgPath != "" {
		config.SetConfigPath(cfgPath)
	}

	gin.DisableConsoleColor()

	adapter := log.GinAdapter()
	gin.DefaultWriter = adapter
	gin.DefaultErrorWriter = adapter

	gin.SetMode(config.ServerCfg().Env)
	http := gin.New()

	http.Use(middleware.JSONLogger("/assets"))
	http.Use(gin.Recovery())

	for _, f := range _waitBrforeRun {
		f()
	}

	for _, f := range _waitInitRoter {
		f(http)
	}

	host := config.ServerCfg().Host
	port := config.ServerCfg().Port
	return http.Run(fmt.Sprintf("%s:%d", host, port))
}
