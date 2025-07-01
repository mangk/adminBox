package adminBox

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
)

var _adminBox *gin.Engine
var _adminBoxInitOnce sync.Once
var _waitInitRoter []func(root *gin.Engine)

func newHttpServer() {
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

func httpEngine() *gin.Engine {
	if _adminBox == nil {
		newHttpServer()
	}
	return _adminBox
}

func SetRouter(f func(root *gin.Engine), setNow ...bool) {
	if len(setNow) > 0 && setNow[0] {
		f(httpEngine())
	} else {
		_waitInitRoter = append(_waitInitRoter, f)
	}
}

func ListenAndServer() {
	addr := Addr()
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

func Run() {
	log.Info("[Project Start]", "listen", Addr())
	defer log.Close()

	for _, f := range _waitInitRoter {
		f(httpEngine())
	}

	httpEngine().Run(Addr())

	log.Info("[Project EXIT]")
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

func Daemon(svcConfig *service.Config) {
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Panicf("[Daemon New Error] %s", err)
		return
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Errorf("[Daemon Control Error] %s", err)
		}
		return
	}

	logger, err := s.Logger(nil)
	if err != nil {
		log.Errorf("[Daemon Logger Error] %s", err)
	}
	err = s.Run()
	if err != nil {
		log.Errorf("[Daemon Run Error] %s", err)
		logger.Error(err)
	}
}

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	log.Info("[Daemon Start]")
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	log.Info("[Daemon run] 1")
	Run()
	log.Info("[Daemon run] 2")
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	log.Info("[Daemon Stop 1]")
	<-time.After(time.Second * 2)

	log.Info("[Daemon Stop 2]")
	return nil
}
