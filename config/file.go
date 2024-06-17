package config

type File struct {
	Name       string  `json:"name" yaml:"name"`
	Driver     string  `json:"driver" yaml:"driver"`         // 类型 cos:腾讯云COS; oss:阿里云OSS; local:本地文件存储；
	Bucket     string  `json:"bucket" yaml:"bucket"`         // 储存桶
	Region     string  `json:"region" yaml:"region"`         // 地域: 腾讯:region; 阿里:Endpoint;
	ID         string  `json:"id" yaml:"id"`                 // id
	Key        string  `json:"key" yaml:"key"`               // key
	PrefixPath string  `json:"prefixPath" yaml:"prefixPath"` // 文件访问前缀
	StorePath  string  `json:"storePath" yaml:"storePath"`   // 本地文件存储路径(类型为local时必填)
	CdnURL     string  `json:"cdnUrl" yaml:"cdnUrl"`         // 文件访问CND域名
	Limit      float64 `json:"limit" yaml:"limit"`
}

type Local struct {
	Path      string `json:"path" yaml:"path"`           // 本地文件访问路径
	StorePath string `json:"storePath" yaml:"storePath"` // 本地文件存储路径
}

type AliyunOSS struct {
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `json:"accessKeyId" yaml:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" yaml:"accessKeySecret"`
	BucketName      string `json:"bucketName" yaml:"bucketName"`
	BucketUrl       string `json:"bucketUrl" yaml:"bucketUrl"`
	BasePath        string `json:"basePath" yaml:"basePath"`
}

type TencentCOS struct {
	Bucket    string `json:"bucket" yaml:"bucket"`       // 储存桶
	Region    string `json:"region" yaml:"region"`       // 地域
	SecretID  string `json:"secretId" yaml:"secretId"`   // id
	SecretKey string `json:"secretKey" yaml:"secretKey"` // key
	BaseURL   string `json:"baseUrl" yaml:"baseUrl"`     // 文件访问CND域名
}

type Qiniu struct {
	Zone          string `json:"zone" yaml:"zone"`                   // 存储区域
	Bucket        string `json:"bucket" yaml:"bucket"`               // 空间名称
	ImgPath       string `json:"imgPath" yaml:"imgPath"`             // CDN加速域名
	UseHTTPS      bool   `json:"useHttps" yaml:"useHttps"`           // 是否使用https
	AccessKey     string `json:"accessKey" yaml:"accessKey"`         // 秘钥AK
	SecretKey     string `json:"secretKey" yaml:"secretKey"`         // 秘钥SK
	UseCdnDomains bool   `json:"useCdnDomains" yaml:"useCdnDomains"` // 上传是否使用CDN上传加速
}

type AwsS3 struct {
	Bucket           string `json:"bucket" yaml:"bucket"`
	Region           string `json:"region" yaml:"region"`
	Endpoint         string `json:"endpoint" yaml:"endpoint"`
	S3ForcePathStyle bool   `json:"s3ForcePathStyle" yaml:"s3ForcePathStyle"`
	DisableSSL       bool   `json:"disableSsl" yaml:"disableSsl"`
	SecretID         string `json:"secretId" yaml:"secretId"`
	SecretKey        string `json:"secretKey" yaml:"secretKey"`
	BaseURL          string `json:"baseUrl" yaml:"baseUrl"`
	PathPrefix       string `json:"pathPrefix" yaml:"pathPrefix"`
}
