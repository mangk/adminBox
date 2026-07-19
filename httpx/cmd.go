package httpx

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

var _cfgFilePath string

type program struct{}

func (p *program) Start(s service.Service) error {
	fmt.Println("▶ 服务 Start 方法被调用")
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	fmt.Println("■ 服务 Stop 方法被调用")
	return nil
}

func (p *program) run() {
	fmt.Println("服务启动[config:" + _cfgFilePath + "]")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go httpServer(_cfgFilePath)

	<-sigChan // 阻塞等待停止信号
	fmt.Println("服务停止")
}

func getService(name, displayName string) service.Service {
	svcConfig := &service.Config{
		Name:        name,
		DisplayName: displayName,
		Description: "",
		UserName:    "",
		Arguments:   []string{"run", "--config", _cfgFilePath}, // service 负责维护注册到系统的启动项，这里设置运行，还是交给了 cobra 去真正运行服务
	}

	s, err := service.New(&program{}, svcConfig)
	if err != nil {
		fmt.Println("创建服务失败:", err)
		os.Exit(1)
	}
	return s
}

func Execute(serverName, serverShort string) {
	// 设置系统命令基本参数
	rootCmd := &cobra.Command{
		Use:   serverName,
		Short: serverShort,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if exe, err := os.Executable(); err != nil {
				return fmt.Errorf("程序运行目录读取错误: %e", err)
			} else {
				exeDir := filepath.Dir(exe)
				if _cfgFilePath != "" && !filepath.IsAbs(_cfgFilePath) {
					_cfgFilePath = filepath.Join(exeDir, _cfgFilePath)
				}
			}

			if _, err := os.Stat(_cfgFilePath); os.IsNotExist(err) {
				return fmt.Errorf("读取配置文件错误文件不存在:[%s]", _cfgFilePath)
			}

			return nil
		},
	}

	rootCmd.PersistentFlags().StringVarP(&_cfgFilePath, "config", "c", "config.yaml", "配置文件路径")

	// 服务管理相关命令
	daemonCmd := &cobra.Command{
		Use:     "daemon",
		Aliases: []string{"d"},
		Short:   "管理系统服务",
		Long:    `安装、启动、停止或卸载系统服务。`,
	}

	// 为服务管理命令增加子命令 install、uninstall、start、stop
	{
		daemonCmd.AddCommand(&cobra.Command{
			Use:   "install",
			Short: "安装服务到系统",
			Run: func(cmd *cobra.Command, args []string) {
				s := getService(serverName, serverShort)
				if err := s.Install(); err != nil {
					fmt.Println("安装失败:", err)
					return
				}
				fmt.Println("✓ 服务安装成功！")
			},
		})

		daemonCmd.AddCommand(&cobra.Command{
			Use:   "uninstall",
			Short: "卸载系统服务",
			Run: func(cmd *cobra.Command, args []string) {
				s := getService(serverName, serverShort)
				if err := s.Uninstall(); err != nil {
					fmt.Println("卸载失败:", err)
					return
				}
				fmt.Println("✓ 服务卸载成功！")
			},
		})

		daemonCmd.AddCommand(&cobra.Command{
			Use:   "start",
			Short: "启动系统服务",
			Run: func(cmd *cobra.Command, args []string) {
				s := getService(serverName, serverShort)
				if err := s.Start(); err != nil {
					fmt.Println("启动失败:", err)
					return
				}
				fmt.Println("✓ 服务启动成功！")
			},
		})

		daemonCmd.AddCommand(&cobra.Command{
			Use:   "stop",
			Short: "停止系统服务",
			Run: func(cmd *cobra.Command, args []string) {
				s := getService(serverName, serverShort)
				if err := s.Stop(); err != nil {
					fmt.Println("停止失败:", err)
					return
				}
				fmt.Println("✓ 服务停止成功！")
			},
		})
	}

	// 将服务管理命令注册到root
	rootCmd.AddCommand()

	// 程序启动命令，kardianos/service 告诉操作系统我要通过 run 命令来运行程序。这里注册 run命令，并负责运行
	rootCmd.AddCommand(
		&cobra.Command{
			Use: "run",
			Run: func(cmd *cobra.Command, args []string) {
				(&program{}).run()
			},
		}, // 运行的命令
		daemonCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
