package config

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var _config *configInstance
var _viper *viper.Viper

type configInstance struct {
	Server server           `json:"server,omitempty" yaml:"server,omitempty"`
	Log    log              `json:"log,omitempty" yaml:"log,omitempty"`
	DB     map[string]DB    `json:"db,omitempty" yaml:"db,omitempty"`
	Cache  map[string]cache `json:"cache,omitempty" yaml:"cache,omitempty"`
}

func (c *configInstance) i() configInstance {
	if _config == nil {
		// TODO 支持从 ENV 读取配置
		cfgFilePath := flag.String("c", "./config.yaml", "config file path")
		flag.Parse()

		_config = &configInstance{}

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
  runAt: 
log:
  prefix: noConfigFile
`)
			if err := _viper.ReadConfig(bytes.NewBuffer(configStr)); err != nil {
				panic(fmt.Errorf("read config error: %s", err))
			}
		}

		if err := _viper.Unmarshal(_config); err != nil {
			panic(fmt.Errorf("read config error: %s", err))
		}
	}

	return *_config
}

func GetAny(path string) any {
	return _viper.Get(path)
}

func Get(path string) string {
	return _viper.GetString(path)
}

func ServerCfg() server {
	return _config.i().Server
}

func CORSCfg() cors {
	return _config.i().Server.CORS
}

func JwtCfg() JWT {
	return _config.i().Server.Jwt
}

func CaptchaCfg() captcha {
	return _config.i().Server.Captcha
}

func LogCfg() log {
	return _config.i().Log
}

func DBCfg() map[string]DB {
	return _config.i().DB
}

func CacheCfg() map[string]cache {
	cfg := *_config
	if cfg.Cache == nil {
		return nil
	}
	return _config.i().Cache
}

func FileCfg() map[string]File {
	return *_config.Server.File
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
