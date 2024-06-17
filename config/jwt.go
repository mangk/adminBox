package config

type JWT struct {
	SigningKey  string `json:"signingKey" yaml:"signingKey"`   // jwt签名
	ExpiresTime string `json:"expiresTime" yaml:"expiresTime"` // 过期时间
	BufferTime  string `json:"bufferTime" yaml:"bufferTime"`   // 缓冲时间
	Issuer      string `json:"issuer" yaml:"issuer"`           // 签发者
}
