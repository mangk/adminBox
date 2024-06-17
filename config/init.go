package config

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/mangk/adminX/moduleRegister"
	"os"

	"github.com/spf13/viper"
)

var _config *config
var _viper *viper.Viper

func init() {
	moduleRegister.ModuleAdd(cfg{})
}

type cfg struct{}

func (cfg) InitModule() {
	cfgFilePath := flag.String("c", "./config.yaml", "config file path")
	flag.Parse()

	_config = &config{}

	_viper = viper.New()
	_viper.SetConfigType("yaml")

	if pathExists(*cfgFilePath) {
		_viper.SetConfigFile(*cfgFilePath)
		if err := _viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("read config file (%s) error: %s", *cfgFilePath, err))
		}
	} else {
		configStr := []byte(`
server:
  name: no config file
  host: 127.0.0.1
  port: 8910
log:
  prefix: no config file
`)
		if err := _viper.ReadConfig(bytes.NewBuffer(configStr)); err != nil {
			panic(fmt.Errorf("read config error: %s", err))
		}
	}

	if err := _viper.Unmarshal(_config); err != nil {
		panic(fmt.Errorf("read config error: %s", err))
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
