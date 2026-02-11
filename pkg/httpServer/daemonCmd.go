package httpServer

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/kardianos/service"
// 	"github.com/mangk/adminBox/pkg/config"
// 	"github.com/mangk/adminBox/pkg/util"
// 	"github.com/spf13/cobra"
// )

// func Execute(serviceFileName, desc string) {
// 	execDir := util.GetExecPath()
// 	os.Chdir(execDir)

// 	var argsConfigFilePath string

// 	rootCmd := &cobra.Command{
// 		Use:   serviceFileName,
// 		Short: desc,
// 		Long:  desc,
// 		Run: func(cmd *cobra.Command, args []string) {
// 			cmd.Help()
// 		},
// 	}
// 	rootCmd.PersistentFlags().StringVarP(&argsConfigFilePath, "config", "c", "", "指定配置文件路径")
// 	if argsConfigFilePath != "" {
// 		config.SetConfigPath(argsConfigFilePath)
// 	}

// 	s, err := service.New(&program{}, &service.Config{
// 		Name:        serviceFileName,
// 		DisplayName: desc,
// 		Description: desc,
// 	})
// 	if err != nil {
// 		panic(fmt.Sprintf("[Daemon Create Error] %s", err))
// 	}

// 	rootCmd.AddCommand(&cobra.Command{
// 		Use:     "serve",
// 		Aliases: []string{"s"},
// 		Short:   "运行程序",
// 		Long:    "直接以阻塞模式运行程序",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			httpServer()
// 		},
// 	})

// 	daemonCmd.AddCommand(&cobra.Command{
// 		Use:   "start",
// 		Short: "启动服务",
// 		Long:  "启动服务，启动服务前，需要通过 install 注册",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			if err := s.Start(); err != nil {
// 				log.Printf("[Daemon Start] err: %s", err)
// 				return
// 			}
// 			log.Printf("[Daemon Start Success]")
// 		},
// 	}, &cobra.Command{
// 		Use:   "stop",
// 		Short: "停止服务",
// 		Long:  "停止正在运行的服务",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			if err := s.Stop(); err != nil {
// 				log.Printf("[Daemon Stop] err: %s", err)
// 				return
// 			}
// 			log.Printf("[Daemon Stop Success]")
// 		},
// 	}, &cobra.Command{
// 		Use:   "restart",
// 		Short: "重启服务",
// 		Long:  "重启服务",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			if err := s.Restart(); err != nil {
// 				log.Printf("[Daemon Restart] err: %s", err)
// 				return
// 			}
// 			log.Printf("[Daemon Restart Success]")
// 		},
// 	}, &cobra.Command{
// 		Use:   "install",
// 		Short: "安装服务",
// 		Long:  "将程序加入到系统的守护进程中，使其能够在后台运行以及跟随系统开机启动",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			if err := s.Install(); err != nil {
// 				log.Printf("[Daemon Install] err: %s", err)
// 				return
// 			}
// 			log.Printf("[Daemon Install Success]")
// 		},
// 	}, &cobra.Command{
// 		Use:   "uninstall",
// 		Short: "卸载服务",
// 		Long:  "将程序从系统服务中移除，以便在系统启动时不再自动启动服务",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			if err := s.Uninstall(); err != nil {
// 				log.Printf("[Daemon Uninstall] err: %s", err)
// 				return
// 			}
// 			log.Printf("[Daemon Uninstall Success]")
// 		},
// 	})

// 	rootCmd.AddCommand(daemonCmd)

// 	if err := rootCmd.Execute(); err != nil {
// 		log.Fatalf("[%s Run Error] %s", serviceFileName, err)
// 		os.Exit(1)
// 	}
// }

// type program struct{}

// func (p *program) Start(s service.Service) error {
// 	log.Print("[Daemon Start]")
// 	go httpServer()
// 	return nil
// }

// func (p *program) Stop(s service.Service) error {
// 	<-time.After(time.Second * 2)
// 	log.Print("[Daemon Stop]")
// 	return nil
// }

import (
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:     "daemon",
	Aliases: []string{"d"},
	Short:   "守护程序运行",
	Long:    "通过守护程序运行",
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "安装服务",
	Long:  "将程序加入到系统的守护进程中，使其能够在后台运行以及跟随系统开机启动",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := newService()
		if err != nil {
			return err
		}
		return s.Install()
	},
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "卸载服务",
	Long:  "将程序从系统服务中移除，以便在系统启动时不再自动启动服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := newService()
		if err != nil {
			return err
		}
		return s.Uninstall()
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务",
	Long:  "启动服务，启动服务前，需要通过 install 注册",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := newService()
		if err != nil {
			return err
		}
		return s.Start()
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止服务",
	Long:  "停止正在运行的服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := newService()
		if err != nil {
			return err
		}
		return s.Stop()
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "重启服务",
	Long:  "重启服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := newService()
		if err != nil {
			return err
		}
		return s.Restart()
	},
}

func init() {
	daemonCmd.AddCommand(
		installCmd,
		uninstallCmd,
		startCmd,
		stopCmd,
		restartCmd,
	)

	rootCmd.AddCommand(daemonCmd)
}
