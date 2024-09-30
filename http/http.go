package http

import (
	"encoding/json"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
	"github.com/mangk/adminBox/moduleRegister"
)

var _adminBox *Http

type Http struct {
	httpClient *gin.Engine
}

func New() *Http {
	_adminBox = &Http{}

	moduleRegister.ModelInit()

	return _adminBox
}

func newHttpServer() {
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
	_adminBox.httpClient = http
}

func HttpEngine() *gin.Engine {
	if _adminBox.httpClient == nil {
		newHttpServer()
	}
	return _adminBox.httpClient
}

func (c *Http) ListenAndServer() {
	addr := Addr()
	log.Info("Project Start", "listen", addr)
	if err := endless.ListenAndServe(addr, c.httpClient); err != nil {
		log.Error("Project Error", "error", err)
	}
	log.Info("Project EXIT")
	defer log.Close()
}

func (c *Http) Run() {
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
