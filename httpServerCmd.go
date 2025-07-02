package adminBox

import (
	"log"
	"os"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

var _configFilePath string
var _userName string
var _Use string
var _Short string
var _Long string

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&_configFilePath, "config", "c", "./config.yaml", "指定配置文件路径")

	rootCmd.AddCommand(daemonStart)

	rootCmd.AddCommand(daemonStop)

	rootCmd.AddCommand(daemonRestart)

	daemonInstall.Flags().StringVarP(&_configFilePath, "config", "c", "", "指定配置文件(绝对)路径")
	// daemonInstall.MarkFlagRequired("config")
	rootCmd.AddCommand(daemonInstall)

	rootCmd.AddCommand(daemonUninstall)

	endlessListenAndServer.Flags().StringVarP(&_configFilePath, "config", "c", "", "指定配置文件路径")
	rootCmd.AddCommand(endlessListenAndServer)
}

func Execute(Use, Short, Long string) {
	rootCmd.Use = Use
	_Use = Use
	rootCmd.Short = Short
	_Short = Short
	rootCmd.Long = Long
	_Long = Long
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var daemonStart = &cobra.Command{
	Use:   "start",
	Short: "启动服务",
	Long:  "启动服务，启动服务前，需要通过 install 注册",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newDaemon(_Use, _Use, _Long, _userName).Start(); err != nil {
			log.Printf("[Daemon Start] err: %s", err)
		}
	},
}

var daemonStop = &cobra.Command{
	Use:   "stop",
	Short: "停止服务",
	Long:  "停止正在运行的服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newDaemon(_Use, _Use, _Long, _userName).Stop(); err != nil {
			log.Printf("[Daemon Stop] err: %s", err)
		}
	},
}

var daemonRestart = &cobra.Command{
	Use:   "restart",
	Short: "重启服务",
	Long:  "重启服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newDaemon(_Use, _Use, _Long, _userName).Restart(); err != nil {
			log.Printf("[Daemon Restart] err: %s", err)
		}
	},
}

var daemonInstall = &cobra.Command{
	Use:   "install",
	Short: "安装服务",
	Long:  "将程序加入到系统的守护进程中，使其能够在后台运行以及跟随系统开机启动",
	Run: func(cmd *cobra.Command, args []string) {
		var s service.Service
		if _configFilePath != "" {
			s = newDaemon(_Use, _Use, _Long, "root", "-c", _configFilePath)
		} else {
			s = newDaemon(_Use, _Use, _Long, _userName)
		}

		if err := s.Install(); err != nil {
			log.Printf("[Daemon Install] err: %s", err)
		}
	},
}

var daemonUninstall = &cobra.Command{
	Use:   "uninstall",
	Short: "卸载服务",
	Long:  "将程序从系统服务中移除，以便在系统启动时不再自动启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := newDaemon(_Use, _Use, _Long, _userName).Uninstall(); err != nil {
			log.Printf("[Daemon Uninstall] err: %s", err)
		}
	},
}

var endlessListenAndServer = &cobra.Command{
	Use:   "endlessRun",
	Short: "平滑重启运行程序",
	Long:  "不将程序注册成服务，但程序以平滑重启方式运行，通过传递信号 kill -1 实现服务平滑重启",
	Run: func(cmd *cobra.Command, args []string) {
		listenAndServer()
	},
}
