package httpServer

import (
	"fmt"
	"os"

	"github.com/mangk/adminBox/pkg/util"
	"github.com/spf13/cobra"
)

var _cfgFilePath string
var _execDir string
var _serverName string
var _serverShort string

var rootCmd = &cobra.Command{}

func Execute(serverName, serverShort string) error {
	_execDir = util.GetExecPath()
	os.Chdir(_execDir)
	fmt.Printf("[HttpServer run pwd %s]\n", _execDir)

	_serverName = serverName
	_serverShort = serverShort

	rootCmd.Use = _serverName
	rootCmd.Short = _serverShort
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&_cfgFilePath,

		"config",
		"c",
		"config.yaml",
		"指定配置文件路径",
	)
}
