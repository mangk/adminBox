package httpx

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
