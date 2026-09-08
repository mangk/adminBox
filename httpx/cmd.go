package httpx

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/kardianos/service"
	"github.com/mangk/adminBox/config"
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

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- httpServer(_cfgFilePath)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-serverErr:
		if err != nil {
			fmt.Println("HTTP服务启动失败:", err)
			os.Exit(1)
		}
	case <-sigChan: // 收到停止信号
		fmt.Println("服务停止")
	}
}

// resolveCfgPath 把相对配置文件路径解析为绝对路径，解析基准是"进程当前工作目录"：
//   - 未提供 -c 时，默认加载 <当前目录>/config.yaml；
//   - 提供 -c 时，按给定路径（相对则相对当前目录）加载。
//
// 注册成系统服务时会把服务的 WorkingDirectory 设为安装时的当前目录，
// 因此服务进程的"当前目录"与前台运行一致，两种方式解析结果相同。
func resolveCfgPath() error {
	if _cfgFilePath == "" {
		return fmt.Errorf("配置文件路径为空")
	}
	if filepath.IsAbs(_cfgFilePath) {
		return nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("读取当前工作目录失败: %w", err)
	}
	_cfgFilePath = filepath.Join(wd, _cfgFilePath)
	return nil
}

// validateCfgPath 解析并校验配置文件真实存在。只有真正要读配置的命令（run、install）需要。
func validateCfgPath() error {
	if err := resolveCfgPath(); err != nil {
		return err
	}
	info, err := os.Stat(_cfgFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("读取配置文件错误, 文件不存在:[%s]", _cfgFilePath)
		}
		return fmt.Errorf("读取配置文件错误:[%s] %w", _cfgFilePath, err)
	}
	if info.IsDir() {
		return fmt.Errorf("读取配置文件错误, 路径是目录:[%s]", _cfgFilePath)
	}
	return nil
}

func getService(name, displayName, workingDir string) service.Service {
	svcConfig := &service.Config{
		Name:        name,
		DisplayName: displayName,
		Description: "",
		UserName:    "",
		// 服务进程的工作目录 = 安装时的当前目录：未用 -c 时服务启动同样按
		// <当前目录>/config.yaml 加载；配置文件内的相对路径也与前台运行保持一致。
		WorkingDirectory: workingDir,
		// install 时 _cfgFilePath 已解析为绝对路径（未用 -c 则为 <当前目录>/config.yaml），
		// 写进 unit（ExecStart=... run --config <绝对路径>），服务进程即使 CWD 异常也能定位。
		Arguments: []string{"run", "--config", _cfgFilePath},
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
	}

	rootCmd.PersistentFlags().StringVarP(&_cfgFilePath, "config", "c", "config.yaml", "配置文件路径（默认取当前目录下的 config.yaml）")

	// 服务管理相关命令
	daemonCmd := &cobra.Command{
		Use:     "daemon",
		Aliases: []string{"d"},
		Short:   "管理系统服务",
		Long:    `安装、启动、停止或卸载系统服务。`,
	}

	daemonCmd.AddCommand(
		&cobra.Command{
			Use:   "install",
			Aliases: []string{"i"},
			Short: "安装服务到系统",
			RunE: func(cmd *cobra.Command, args []string) error {
				// 未提供 -c 时校验 <当前目录>/config.yaml 是否存在，避免装上一个启动即失败的服务
				if err := validateCfgPath(); err != nil {
					return err
				}
				wd, err := os.Getwd()
				if err != nil {
					return fmt.Errorf("读取当前工作目录失败: %w", err)
				}
				s := getService(serverName, serverShort, wd)
				if err := s.Install(); err != nil {
					return fmt.Errorf("安装失败: %w", err)
				}
				fmt.Println("✓ 服务安装成功！")
				return nil
			},
		},

		&cobra.Command{
			Use:   "uninstall",
			Aliases: []string{"u"},
			Short: "卸载系统服务",
			RunE: func(cmd *cobra.Command, args []string) error {
				s := getService(serverName, serverShort, "")
				if err := s.Uninstall(); err != nil {
					return fmt.Errorf("卸载失败: %w", err)
				}
				fmt.Println("✓ 服务卸载成功！")
				return nil
			},
		},

		&cobra.Command{
			Use:   "start",
			Aliases: []string{"s"},
			Short: "启动系统服务",
			RunE: func(cmd *cobra.Command, args []string) error {
				s := getService(serverName, serverShort, "")
				if err := s.Start(); err != nil {
					return fmt.Errorf("启动失败: %w", err)
				}
				fmt.Println("✓ 服务启动成功！")
				return nil
			},
		},

		&cobra.Command{
			Use:   "stop",
			Aliases: []string{"p"},
			Short: "停止系统服务",
			RunE: func(cmd *cobra.Command, args []string) error {
				s := getService(serverName, serverShort, "")
				if err := s.Stop(); err != nil {
					return fmt.Errorf("停止失败: %w", err)
				}
				fmt.Println("✓ 服务停止成功！")
				return nil
			},
		},
	)

	// 程序启动命令。systemd 通过 unit 里的 "run --config <绝对路径>" 参数进入这里。
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "运行服务（前台；也是系统服务实际启动入口）",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validateCfgPath(); err != nil {
				return err
			}
			// 在任何异步启动、日志/DB 初始化之前把配置路径告知 config 包
			config.SetConfigPath(_cfgFilePath)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			(&program{}).run()
			return nil
		},
	}

	rootCmd.AddCommand(runCmd, daemonCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
