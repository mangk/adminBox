package config

type captcha struct {
	KeyLong            int `json:"keyLong" yaml:"keyLong"`                       // 验证码长度
	ImgWidth           int `json:"imgWidth" yaml:"imgWidth"`                     // 验证码宽度
	ImgHeight          int `json:"imgHeight" yaml:"imgHeight"`                   // 验证码高度
	Overtime           int `json:"overtime" yaml:"overtime"`                     // 超时时间
}
