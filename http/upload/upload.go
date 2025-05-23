package upload

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
	"io"
	"mime/multipart"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader, keyPrefix ...string) (reqPath, fileKey, md5 string, err error)
	DeleteFile(key string) error
}

func NewOss(ossType ...string) OSS {
	t := "default"
	if len(ossType) > 0 {
		t = ossType[0]
	}
	cfg, ok := config.FileCfg()[t]

	if ok {
		switch cfg.Driver {
		case "local":
			return &Local{cfg: cfg}
		case "cos":
			return &TencentCOS{cfg: cfg}
		case "oss":
			return &AliyunOSS{cfg: cfg}
		case "qiniu":
			return &Qiniu{cfg: cfg}
		//case "aws-s3":
		//	return &AwsS3{}
		default:
			log.Zaplog().Panic(fmt.Sprintf("[%s]未知的对象存储", t))
			return nil
		}
	}
	log.Zaplog().Panic(fmt.Sprintf("[%s]未知的对象存储", t))
	return nil
}

func fileMd5(file multipart.File) string {
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	file.Seek(0, 0)
	return hex.EncodeToString(hash.Sum(nil))
}
