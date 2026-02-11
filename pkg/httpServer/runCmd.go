package httpServer

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "运行程序",
	Long:  "直接以阻塞模式运行程序",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := newService()
		if err != nil {
			return err
		}
		return s.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
