package adminBox

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mangk/adminBox/config"
	"github.com/spf13/cobra"
)

var (
	_userName        string // 执行程序的用户名
	_serviceFileName string // 用来 daemon 创建服务的文件名
	_desc            string // 程序描述
)

var (
	_execDir string // 可执行程序的所在绝对目录
)
var (
	_argsConfigFilePath string // 来自 args 的配置文件地址
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// panic("TODO")
	},
}

func init() {
	cobra.OnInitialize(initExecPwd)
	rootCmd.PersistentFlags().StringVarP(&_argsConfigFilePath, "config", "c", "config.yaml", "指定配置文件路径")

	rootCmd.AddCommand(serve)
	// rootCmd.AddCommand(serverEndless)

	rootCmd.AddCommand(daemon)
	daemon.AddCommand(daemonStart)
	daemon.AddCommand(daemonStop)
	daemon.AddCommand(daemonRestart)
	daemon.AddCommand(daemonInstall)
	daemon.AddCommand(daemonUninstall)
}

func initExecPwd() {
	pwd, _ := os.Getwd()
	args0 := os.Args[0]

	if filepath.Dir(args0) == pwd {
		_execDir = pwd
	} else {
		if pwd == "/" {
			_execDir = filepath.Dir(args0)
		} else {
			if strings.Contains(args0, os.TempDir()) || (strings.Contains(args0, "Caches") && strings.Contains(args0, "go-build")) {
				_execDir = pwd
			} else {
				_execDir = filepath.Dir(filepath.Join(pwd, args0))
			}
		}
	}

	os.Chdir(_execDir)
	println("程序目录: " + _execDir)
}

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
	_serviceFileName = serviceFileName
	_desc = desc

	rootCmd.Use = _serviceFileName
	rootCmd.Short = _desc
	rootCmd.Long = _desc

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("[%s Run Error] %s", _serviceFileName, err)
		os.Exit(1)
	}
}

var serve = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "运行程序",
	Long:    "直接以阻塞模式运行程序",
	Run: func(cmd *cobra.Command, args []string) {
		config.SetConfigPath(buildConfigFilePath())
		run()
	},
}

var daemon = &cobra.Command{
	Use:     "daemon",
	Aliases: []string{"d"},
	Short:   "守护程序运行",
	Long:    "通过守护程序运行",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var daemonStart = &cobra.Command{
	Use:   "start",
	Short: "启动服务",
	Long:  "启动服务，启动服务前，需要通过 install 注册",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newDaemon(_serviceFileName, _desc, _desc, _userName).Start(); err != nil {
			log.Printf("[Daemon Start] err: %s", err)
			return
		}
		log.Printf("[Daemon Start Success]")
	},
}

var daemonStop = &cobra.Command{
	Use:   "stop",
	Short: "停止服务",
	Long:  "停止正在运行的服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newDaemon(_serviceFileName, _desc, _desc, _userName).Stop(); err != nil {
			log.Printf("[Daemon Stop] err: %s", err)
			return
		}
		log.Printf("[Daemon Stop Success]")
	},
}

var daemonRestart = &cobra.Command{
	Use:   "restart",
	Short: "重启服务",
	Long:  "重启服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newDaemon(_serviceFileName, _desc, _desc, _userName).Restart(); err != nil {
			log.Printf("[Daemon Restart] err: %s", err)
			return
		}
		log.Printf("[Daemon Restart Success]")
	},
}

var daemonInstall = &cobra.Command{
	Use:   "install",
	Short: "安装服务",
	Long:  "将程序加入到系统的守护进程中，使其能够在后台运行以及跟随系统开机启动",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newDaemon(_serviceFileName, _desc, _desc, _userName, "s", "-c", buildConfigFilePath()).Install(); err != nil {
			log.Printf("[Daemon Install] err: %s", err)
			return
		}
		log.Printf("[Daemon Install Success]")
	},
}

var daemonUninstall = &cobra.Command{
	Use:   "uninstall",
	Short: "卸载服务",
	Long:  "将程序从系统服务中移除，以便在系统启动时不再自动启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newDaemon(_serviceFileName, _desc, _desc, _userName).Uninstall(); err != nil {
			log.Printf("[Daemon Uninstall] err: %s", err)
			return
		}
		log.Printf("[Daemon Uninstall Success]")
	},
}
