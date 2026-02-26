package daemon

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kardianos/service"
	plog "github.com/mangk/adminBox/pkg/log"
	"github.com/mangk/adminBox/pkg/util"
	"github.com/spf13/cobra"
)

func buildConfigFilePath() string {
	if _argsConfigFilePath == "" {
		_argsConfigFilePath = "config.yaml"
	}

	var path string
	if []rune(_argsConfigFilePath)[0] != '/' {
		path = filepath.Join(_execDir, _argsConfigFilePath)
	} else {
		path = _argsConfigFilePath
	}

	_, err := os.Stat(path)
	if err != nil {
		panic("配置文件路径错误，请提供正确的配置文件路径（建议使用绝对路径）")
	}
	if os.IsNotExist(err) {
		panic("配置文件路径错误，请提供正确的配置文件路径（建议使用绝对路径）")
	}

	return path
}

func Execute(serviceFileName, desc string) {
	execDir := util.GetExecPath()
	os.Chdir(execDir)

	var argsConfigFilePath string

	program := &program{}
	s, err := service.New(program, &service.Config{
		Name:        serviceFileName,
		DisplayName: desc,
		Description: desc,
		Arguments:   []string{"s", "-c", buildConfigFilePath()},
	})
	if err != nil {
		panic(fmt.Sprintf("[Daemon Create Error] %s", err))
	}

	rootCmd := &cobra.Command{
		Use:   serviceFileName,
		Short: desc,
		Long:  desc,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.PersistentFlags().StringVarP(&argsConfigFilePath, "config", "c", "config.yaml", "指定配置文件路径")

	rootCmd.AddCommand(&cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "运行程序",
		Long:    "直接以阻塞模式运行程序",
		Run: func(cmd *cobra.Command, args []string) {
			// config.SetConfigPath(buildConfigFilePath())
			program.run()
		},
	})

	daemonCmd := &cobra.Command{
		Use:     "daemon",
		Aliases: []string{"d"},
		Short:   "守护程序运行",
		Long:    "通过守护程序运行",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	daemonCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "启动服务",
		Long:  "启动服务，启动服务前，需要通过 install 注册",
		Run: func(cmd *cobra.Command, args []string) {
			if err := s.Start(); err != nil {
				log.Printf("[Daemon Start] err: %s", err)
				return
			}
			log.Printf("[Daemon Start Success]")
		},
	}, &cobra.Command{
		Use:   "stop",
		Short: "停止服务",
		Long:  "停止正在运行的服务",
		Run: func(cmd *cobra.Command, args []string) {
			if err := s.Stop(); err != nil {
				log.Printf("[Daemon Stop] err: %s", err)
				return
			}
			log.Printf("[Daemon Stop Success]")
		},
	}, &cobra.Command{
		Use:   "restart",
		Short: "重启服务",
		Long:  "重启服务",
		Run: func(cmd *cobra.Command, args []string) {
			if err := s.Restart(); err != nil {
				log.Printf("[Daemon Restart] err: %s", err)
				return
			}
			log.Printf("[Daemon Restart Success]")
		},
	}, &cobra.Command{
		Use:   "install",
		Short: "安装服务",
		Long:  "将程序加入到系统的守护进程中，使其能够在后台运行以及跟随系统开机启动",
		Run: func(cmd *cobra.Command, args []string) {
			if err := s.Install(); err != nil {
				log.Printf("[Daemon Install] err: %s", err)
				return
			}
			log.Printf("[Daemon Install Success]")
		},
	}, &cobra.Command{
		Use:   "uninstall",
		Short: "卸载服务",
		Long:  "将程序从系统服务中移除，以便在系统启动时不再自动启动服务",
		Run: func(cmd *cobra.Command, args []string) {
			if err := s.Uninstall(); err != nil {
				log.Printf("[Daemon Uninstall] err: %s", err)
				return
			}
			log.Printf("[Daemon Uninstall Success]")
		},
	})

	rootCmd.AddCommand(daemonCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("[%s Run Error] %s", serviceFileName, err)
		os.Exit(1)
	}
}

type program struct{}

func (p *program) Start(s service.Service) error {
	plog.Info("[Daemon Start]")
	go p.run()
	return nil
}
func (p *program) run() {
	plog.Info("[Daemon run]")
	plog.Info("[Project Start]", "listen", GetServerAddr())
	defer plog.Close()

	for _, f := range _waitBrforeRun {
		f()
	}

	for _, f := range _waitInitRoter {
		f(httpEngine())
	}

	httpEngine().Run(GetServerAddr())

	log.Info("[Project EXIT]")
}
func (p *program) Stop(s service.Service) error {
	<-time.After(time.Second * 2)
	plog.Info("[Daemon Stop]")
	return nil
}
