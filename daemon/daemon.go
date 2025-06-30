package daemon

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mangk/adminBox"

	"github.com/kardianos/service"
)

func Daemon(svcConfig *service.Config) {
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal("Cannot create service:", err)
	}

	logger, err := s.Logger(nil)
	if err == nil {
		log.SetOutput(os.Stdout)
	}

	// Run service
	if err := s.Run(); err != nil {
		logger.Error(err)
	}

}

type program struct {
	exit    chan struct{}
	signals chan os.Signal
}

func (p *program) Start(s service.Service) error {
	log.Println("[Service] Starting")
	p.exit = make(chan struct{})
	p.signals = make(chan os.Signal, 1)

	go p.run()
	go p.handleSignal()

	return nil
}

func (p *program) run() {
	adminBox.ListenAndServer()
}

func (p *program) Stop(s service.Service) error {
	log.Println("[Service] Stopping")
	close(p.exit)
	return nil
}

func (p *program) handleSignal() {
	signal.Notify(p.signals, syscall.SIGHUP)

	for {
		select {
		case sig := <-p.signals:
			if sig == syscall.SIGHUP {
				log.Println("[Signal] Received SIGHUP – Reloading config (simulate)")
				// Simulate config reload or log rotation, etc.
			}
		case <-p.exit:
			log.Println("[Signal] Exiting signal handler")
			return
		}
	}
}
