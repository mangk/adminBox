package adminBox

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
	"github.com/mangk/adminBox/util"
)

var _adminBox *gin.Engine
var _adminBoxInitOnce sync.Once
var _waitInitRoter []func(root *gin.Engine)

func SetRouter(f func(root *gin.Engine), setNow ...bool) {
	if len(setNow) > 0 && setNow[0] {
		f(httpEngine())
	} else {
		_waitInitRoter = append(_waitInitRoter, f)
	}
}

func GetServerAddr() string {
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

func httpEngine() *gin.Engine {
	if _adminBox == nil {
		_adminBoxInitOnce.Do(func() {
			gin.DisableConsoleColor()
			gin.DefaultWriter = log.GinAdapter() // 设置日志输出到 zaplog
			gin.SetMode(config.ServerCfg().Env)
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

			_adminBox = http
			_waitInitRoter = make([]func(root *gin.Engine), 0)
		})
	}
	return _adminBox
}
func listenAndServer() {
	os.Chdir(util.GetExecPath())
	addr := GetServerAddr()
	log.Info("[Project Start]", "listen", addr)
	defer log.Close()

	for _, f := range _waitInitRoter {
		f(httpEngine())
	}

	es := endless.NewServer(addr, httpEngine())
	es.BeforeBegin = func(add string) {}

	if err := es.ListenAndServe(); err != nil {
		log.Error("[Project Error]", "error", err)
	}

	log.Info("[Project EXIT]")
}

func run() {
	os.Chdir(util.GetExecPath())
	log.Info("[Project Start]", "listen", GetServerAddr())
	defer log.Close()

	for _, f := range _waitInitRoter {
		f(httpEngine())
	}

	httpEngine().Run(GetServerAddr())

	log.Info("[Project EXIT]")
}
