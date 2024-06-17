package http

import (
	"encoding/json"
	"fmt"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/mangk/gAdmin/config"
	"github.com/mangk/gAdmin/log"
)

var _gAdmin *Core

type Core struct {
	httpClient *gin.Engine
}

func New() *Core {
	if _gAdmin == nil {
		gin.DisableConsoleColor()
		gin.DefaultWriter = log.GinAdapter() // 设置日志输出到 zaplog

		http := gin.New()
		http.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			m := []interface{}{}
			m = append(m, "status", fmt.Sprintf("%3d", param.StatusCode),
				"latency", fmt.Sprintf("%v", param.Latency),
				"clientIP", param.ClientIP,
				"method", param.Method,
				"path", param.Path,
				"errorMessage", param.ErrorMessage,
			)
			b, _ := json.Marshal(m)
			return string(b)
		}))
		http.Use(gin.Recovery())

		_gAdmin = &Core{httpClient: http}

	}
	return _gAdmin
}

func HttpEngine() *gin.Engine {
	return _gAdmin.httpClient
}

func (c *Core) ListenAdnServer() {
	addr := Addr()
	log.Info("Project Start", "listen", addr)
	if err := endless.ListenAndServe(addr, c.httpClient); err != nil {
		log.Error("Project Error", "error", err)
	}
	log.Info("Project EXIT")
	defer log.Close()
}

func (c *Core) Run() {
	log.Info("Project Start", "listen", Addr())
	c.httpClient.Run(Addr())
	log.Info("Project EXIT")
	defer log.Close()
}

func Addr() string {
	var host string
	var port int
	if config.ServerCfg().Host != "" {
		host = config.ServerCfg().Host
	}
	if config.ServerCfg().Port != 0 {
		port = config.ServerCfg().Port
	}
	return fmt.Sprintf("%s:%d", host, port)
}
