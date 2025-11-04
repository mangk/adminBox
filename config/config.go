package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/mangk/adminBox/util"
	"github.com/spf13/viper"
)

var _config *configInstance
var _viper *viper.Viper
var _configInitOnce sync.Once
var _configPath string

type configInstance struct {
	Server server           `json:"server,omitempty" yaml:"server,omitempty"`
	Log    log              `json:"log,omitempty" yaml:"log,omitempty"`
	DB     map[string]DB    `json:"db,omitempty" yaml:"db,omitempty"`
	Cache  map[string]cache `json:"cache,omitempty" yaml:"cache,omitempty"`
}

func SetConfigPath(path string) {
	_configPath = path
}

func GetConfigPath() string {
	return _configPath
}

func (c *configInstance) i() configInstance {
	_configInitOnce.Do(func() {
		if _configPath == "" { // 指定路径优先级最高
			_configPath = filepath.Join(util.GetExecPath(), "config.yaml")
			SetConfigPath(_configPath)
			os.Chdir(_configPath)
		}

		_config = &configInstance{}

		_viper = viper.New()
		_viper.SetConfigType("yaml")
		if pathExists(_configPath) {
			_viper.SetConfigFile(_configPath)
			if err := _viper.ReadInConfig(); err != nil {
				panic(fmt.Errorf("read config file (%s) error: %s", _configPath, err))
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
  output:
    - console
`)
			println("USE DEFAULT CONFIG!")
			println(string(configStr))
			if err := _viper.ReadConfig(bytes.NewBuffer(configStr)); err != nil {
				panic(fmt.Errorf("read config error: %s", err))
			}
		}

		if err := _viper.Unmarshal(_config); err != nil {
			panic(fmt.Errorf("read config error: %s", err))
		}
	})

	return *_config
}

func Get(path string) any {
	if _viper ==nil {
		_config.i()
	}
	return _viper.Get(path)
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
	c := _config.i()
	if c.Server.File != nil {
		return *c.Server.File
	}
	return map[string]File{}
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
