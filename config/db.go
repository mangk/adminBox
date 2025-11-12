package config

type DB struct {
	Driver       string `json:"driver" yaml:"driver"`             // 驱动
	Path         string `json:"path" yaml:"path"`                 // 服务器地址
	Port         string `json:"port" yaml:"port"`                 //端口
	Config       string `json:"config" yaml:"config"`             // 高级配置
	Dbname       string `json:"dbname" yaml:"dbname"`             // 数据库名
	Username     string `json:"username" yaml:"username"`         // 数据库用户名
	Password     string `json:"password" yaml:"password"`         // 数据库密码
	Prefix       string `json:"prefix" yaml:"prefix"`             //全局表前缀，单独定义TableName则不生效
	Singular     bool   `json:"singular" yaml:"singular"`         //是否开启全局禁用复数，true表示开启
	MaxIdleConns int    `json:"maxIdleConns" yaml:"maxIdleConns"` // 空闲中的最大连接数
	MaxOpenConns int    `json:"maxOpenConns" yaml:"maxOpenConns"` // 打开到数据库的最大连接数
	LogMode      int    `json:"logMode" yaml:"logMode"`           // 日志等级 Silent:1;Error:2;Warn:3;Info:4
}
