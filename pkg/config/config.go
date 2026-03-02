package config

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var _config *configInstance
var _viper *viper.Viper
var configInit sync.Once
var configPathSet sync.Once
var _configPath = "./config.yaml"

type configInstance struct {
	Server server           `json:"server,omitempty" yaml:"server,omitempty"`
	Log    log              `json:"log,omitempty" yaml:"log,omitempty"`
	DB     map[string]DB    `json:"db,omitempty" yaml:"db,omitempty"`
	Cache  map[string]cache `json:"cache,omitempty" yaml:"cache,omitempty"`
}

func i() *configInstance {
	configInit.Do(func() {
		_viper = viper.New()
		_viper.SetConfigFile(_configPath)
		_viper.SetConfigType("yaml")
		err := _viper.ReadInConfig()
		fmt.Printf("config path set to %s\n", _configPath)

		if err != nil {
			// If config file is not found, use a default one
			fmt.Println("config file not found, using default settings.")
			defaultConfig := []byte(`
server:
  name: adminBox
  port: 8910
  host: 0.0.0.0
log:
  prefix: adminBox
  output:
    - console
`)
			err = _viper.ReadConfig(bytes.NewBuffer(defaultConfig))
		}

		if err == nil {
			_config = &configInstance{}
			err = _viper.Unmarshal(_config)
			if err != nil {
				fmt.Printf("[Unmarshal Error] %s", err)
				os.Exit(1)
			}
		}
	})
	return _config
}

func SetConfigPath(path string) {
	configPathSet.Do(func() {
		_configPath = path
	})
}

func Get(path string) any {
	if _viper == nil {
		i()
	}
	return _viper.Get(path)
}

func ServerCfg() server {
	return i().Server
}

func CORSCfg() cors {
	return i().Server.CORS
}

func JwtCfg() JWT {
	return i().Server.Jwt
}

func CaptchaCfg() captcha {
	return i().Server.Captcha
}

func LogCfg() log {
	return i().Log
}

func DBCfg() map[string]DB {
	return i().DB
}

func CacheCfg() map[string]cache {
	cfg := i()
	if cfg.Cache == nil {
		return nil
	}
	return cfg.Cache
}

func FileCfg() map[string]File {
	c := i()
	if c.Server.File != nil {
		return *c.Server.File
	}
	return map[string]File{}
}

func Custom[T any]() (T, error) {
	var c T
	if _viper == nil {
		i()
	}
	return c, _viper.Unmarshal(&c)
}
