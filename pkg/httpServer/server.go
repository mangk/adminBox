package httpServer

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/pkg/config"
)

var _waitInitRoter = make([]func(root *gin.Engine), 0)
var _waitBrforeRun = make([]func(), 0)

func SetRouter(f func(root *gin.Engine)) {
	_waitInitRoter = append(_waitInitRoter, f)
}

func SetBeforeRun(f func()) {
	_waitBrforeRun = append(_waitBrforeRun, f)
}

func httpServer() {
	gin.DisableConsoleColor()

	// adapter := log.GinAdapter()
	// gin.DefaultWriter = adapter
	// gin.DefaultErrorWriter = adapter

	// gin.SetMode(config.ServerCfg().Env)
	http := gin.New()
	http.Use(gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		m := map[string]interface{}{
			"status":       p.StatusCode,
			"latency":      p.Latency.String(),
			"clientIP":     p.ClientIP,
			"method":       p.Method,
			"path":         p.Path,
			"errorMessage": p.ErrorMessage,
		}
		b, _ := json.Marshal(m)
		return string(b)
	}))

	http.Use(gin.Recovery())

	for _, f := range _waitBrforeRun {
		f()
	}

	for _, f := range _waitInitRoter {
		f(http)
	}

	host := config.ServerCfg().Host
	port := config.ServerCfg().Port
	http.Run(fmt.Sprintf("%s:%d", host, port))
}
