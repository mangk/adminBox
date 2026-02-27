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
var _configInitOnce sync.Once

type configInstance struct {
	Server server           `json:"server,omitempty" yaml:"server,omitempty"`
	Log    log              `json:"log,omitempty" yaml:"log,omitempty"`
	DB     map[string]DB    `json:"db,omitempty" yaml:"db,omitempty"`
	Cache  map[string]cache `json:"cache,omitempty" yaml:"cache,omitempty"`
}

// Load initializes the configuration from a given path.
// This function should be called once from the main application.
func Load(path string, v interface{}) error {
	var err error
	_configInitOnce.Do(func() {
		_viper = viper.New()
		_viper.SetConfigFile(path)
		_viper.SetConfigType("yaml")

		if err = _viper.ReadInConfig(); err != nil {
			// If config file is not found, use a default one
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Println("config file not found, using default settings.")
				defaultConfig := []byte(`
server:
  name: adminBox
  port: ":8080"
log:
  prefix: adminBox
  output:
    - console
`)
				err = _viper.ReadConfig(bytes.NewBuffer(defaultConfig))
			}
		}

		if err == nil {
			_config = &configInstance{}
			err = _viper.Unmarshal(_config)
		}

		if err == nil && v != nil {
			err = _viper.Unmarshal(v)
		}
	})

	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	return nil
}

func i() *configInstance {
	if _config == nil {
		// This will panic if Load() has not been called.
		// This is intentional to enforce proper initialization.
		panic("configuration has not been loaded. Please call config.Load() in your main function.")
	}
	return _config
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

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
