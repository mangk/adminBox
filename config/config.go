package config

type config struct {
	Server *server           `json:"server,omitempty" yaml:"server,omitempty"`
	Log    *log              `json:"log,omitempty" yaml:"log,omitempty"`
	DB     *map[string]DB    `json:"db,omitempty" yaml:"db,omitempty"`
	Cache  *map[string]cache `json:"cache,omitempty" yaml:"cache,omitempty"`
}

func GetAny(path string) any {
	return _viper.Get(path)
}

func Get(path string) string {
	return _viper.GetString(path)
}

func ServerCfg() server {
	return *_config.Server
}

func CORSCfg() cors {
	return _config.Server.CORS
}

func JwtCfg() JWT {
	return _config.Server.Jwt
}

func CaptchaCfg() captcha {
	return _config.Server.Captcha
}

func LogCfg() log {
	return *_config.Log
}

func DBCfg() map[string]DB {
	return *_config.DB
}

func CacheCfg() map[string]cache {
	cfg := *_config
	if cfg.Cache == nil {
		return nil
	}
	return *_config.Cache
}

func FileCfg() map[string]File {
	return *_config.Server.File
}
