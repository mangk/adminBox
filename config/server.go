package config

type server struct {
	Name              string `json:"name" yaml:"name"`
	Env               string `json:"env" yaml:"env"`                             // 运行环境
	Host              string `json:"host" yaml:"host"`                           // 监听地址
	Port              int    `json:"port" yaml:"port"`                           // 端口值
	UseMultipoint     bool   `json:"useMultipoint" yaml:"useMultipoint"`         // 多点登录拦截
	FrontHost         string `json:"frontHost" yaml:"frontHost"`                 // 前后端分离部署时，指向前端主页
	FrontRouterPrefix string `json:"frontRouterPrefix" yaml:"frontRouterPrefix"` // 前端项目前缀
	Endless           bool   `json:"endless" yaml:"endless"`                     // 通过endless启动

	CORS    cors             `json:"cors,omitempty" yaml:"cors,omitempty"`
	Jwt     jwt              `json:"jwt,omitempty" yaml:"jwt,omitempty"`
	Captcha captcha          `json:"captcha,omitempty" yaml:"captcha,omitempty"`
	File    *map[string]file `json:"file,omitempty" yaml:"file,omitempty"`
}
