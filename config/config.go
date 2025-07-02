package config

import (
	"bytes"
	"flag"
	"fmt"
	defaultLog "log"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

var _config *configInstance
var _viper *viper.Viper
var _configInitOnce sync.Once

type configInstance struct {
	Server server           `json:"server,omitempty" yaml:"server,omitempty"`
	Log    log              `json:"log,omitempty" yaml:"log,omitempty"`
	DB     map[string]DB    `json:"db,omitempty" yaml:"db,omitempty"`
	Cache  map[string]cache `json:"cache,omitempty" yaml:"cache,omitempty"`
}

func getConfigPath() string {
	cfgFilePath := flag.String("c", "", "config file path")
	flag.Parse()

	if *cfgFilePath != "" {
		if _, err := os.Stat(*cfgFilePath); err == nil {
			return *cfgFilePath
		} else {
			defaultLog.Fatalf("指定的配置文件路径不存在: %s", *cfgFilePath)
		}
	}

	// fallback: 使用可执行文件所在目录
	exePath, err := os.Executable()
	if err != nil {
		defaultLog.Fatalf("获取可执行文件路径失败: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	defaultCfg := filepath.Join(exeDir, "config.yaml")

	if _, err := os.Stat(defaultCfg); err == nil {
		return defaultCfg
	}

	defaultLog.Fatalf("未提供配置文件路径，且在可执行文件目录下未找到 config.yaml: %s", defaultCfg)
	return ""
}

func (c *configInstance) i() configInstance {
	_configInitOnce.Do(func() {
		configPath := getConfigPath()

		_config = &configInstance{}

		_viper = viper.New()
		_viper.SetConfigType("yaml")
		if pathExists(configPath) {
			_viper.SetConfigFile(configPath)
			if err := _viper.ReadInConfig(); err != nil {
				panic(fmt.Errorf("read config file (%s) error: %s", configPath, err))
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
	})

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
	return *_config.i().Server.File
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
