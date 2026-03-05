package httpServer

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/pkg/config"
	"github.com/mangk/adminBox/pkg/log"
	"github.com/mangk/adminBox/pkg/middleware"

	"context"

	"github.com/kardianos/service"
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

type program struct {
	ctx     context.Context
	cancel  context.CancelFunc
	cfgPath string
}

func (p *program) Start(s service.Service) error {
	p.ctx, p.cancel = context.WithCancel(context.Background())
	go func() {
		if err := httpServer(p.cfgPath); err != nil {
			fmt.Printf("Server failed: %v\n", err)
			os.Exit(1)
		}
	}()
	return nil
}

func (p *program) Stop(s service.Service) error {
	if p.cancel != nil {
		p.cancel()
	}
	return nil
}

func newService() (service.Service, error) {
	return service.New(&program{cfgPath: _cfgFilePath}, &service.Config{
		Name:        _serverName,
		DisplayName: _serverName,
		Description: _serverShort,
		Arguments:   []string{"run", "-c", _cfgFilePath},
	})
}
