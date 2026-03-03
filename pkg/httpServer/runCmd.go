package httpServer

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "运行程序",
	Long:  "直接以阻塞模式运行程序",
	RunE: func(cmd *cobra.Command, args []string) error {
		println(cfgFilePath, 2222)
		httpServer(cfgFilePath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
