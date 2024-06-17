package config

type captcha struct {
	KeyLong            int `json:"keyLong" yaml:"keyLong"`                       // 验证码长度
	ImgWidth           int `json:"imgWidth" yaml:"imgWidth"`                     // 验证码宽度
	ImgHeight          int `json:"imgHeight" yaml:"imgHeight"`                   // 验证码高度
	OpenCaptcha        int `json:"openCaptcha" yaml:"openCaptcha"`               // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int `json:"openCaptchaTimeout" yaml:"openCaptchaTimeout"` // 防爆破验证码超时时间，单位：s(秒)
}
