package config

type server struct {
	Name                string `json:"name" yaml:"name"`
	Desc                string `json:"desc" yaml:"desc"`
	Logo                string `json:"logo" yaml:"logo"`
	Env                 string `json:"env" yaml:"env"`                                 // 运行环境
	Host                string `json:"host" yaml:"host"`                               // 监听地址
	Port                int    `json:"port" yaml:"port"`                               // 端口值
	UseMultipoint       bool   `json:"useMultipoint" yaml:"useMultipoint"`             // 多点登录拦截
	FrontHost           string `json:"frontHost" yaml:"frontHost"`                     // 前后端分离部署时，指向前端主页
	FrontRouterPrefix   string `json:"frontRouterPrefix" yaml:"frontRouterPrefix"`     // 前端项目前缀
	BackendRouterPrefix string `json:"backendRouterPrefix" yaml:"backendRouterPrefix"` // 后端项目接口前缀
	RunAt               string `json:"runAt" yaml:"runAt"`                             // 运行位置
	Endless             bool   `json:"endless" yaml:"endless"`                         // 通过endless启动

	CORS    cors             `json:"cors,omitempty" yaml:"cors,omitempty"`
	Jwt     JWT              `json:"jwt,omitempty" yaml:"jwt,omitempty"`
	Captcha captcha          `json:"captcha,omitempty" yaml:"captcha,omitempty"`
	File    *map[string]File `json:"file,omitempty" yaml:"file,omitempty"`
}
