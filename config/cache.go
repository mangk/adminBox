package config

type cache struct {
	DB       int    `json:"db" yaml:"db"`             // 数据库index
	Addr     string `json:"addr" yaml:"addr"`         // 服务器地址:端口
	Password string `json:"password" yaml:"password"` // 密码
}
