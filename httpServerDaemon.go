package adminBox

import (
	"fmt"
	"time"

	"github.com/kardianos/service"
	"github.com/mangk/adminBox/log"
)

func newDaemon(Name, DisplayName, Description, UserName string, args ...string) service.Service {
	prg := &program{}
	cfg := &service.Config{
		Name:        Name,
		DisplayName: DisplayName,
		Description: Description,
	}
	if UserName != "" {
		cfg.UserName = UserName
	}
	if len(args) > 0 {
		cfg.Arguments = args
	}
	s, err := service.New(prg, cfg)
	if err != nil {
		panic(fmt.Sprintf("[Daemon Create Error] %s", err))
	}
	return s
}

type program struct{}

func (p *program) Start(s service.Service) error {
	log.Info("[Daemon Start]")
	go p.run()
	return nil
}
func (p *program) run() {
	log.Info("[Daemon run]")
	run()
}
func (p *program) Stop(s service.Service) error {
	<-time.After(time.Second * 2)
	log.Info("[Daemon Stop]")
	return nil
}
