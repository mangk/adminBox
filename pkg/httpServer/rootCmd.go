package httpServer

import (
	"os"

	"github.com/mangk/adminBox/pkg/util"
	"github.com/spf13/cobra"
)

var cfgFilePath string
var _serverName string
var _serverShort string

var rootCmd = &cobra.Command{}

func Execute(serverName, serverShort string) error {
	dir := util.GetExecPath()
	os.Chdir(dir)
	println(dir)

	_serverName = serverName
	_serverShort = serverShort

	rootCmd.Use = _serverName
	rootCmd.Short = _serverShort
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&cfgFilePath,

		"config",
		"c",
		"config.yaml",
		"指定配置文件路径",
	)
}
