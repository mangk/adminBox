package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"

	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

type TencentCOS struct {
	cfg config.File
}

func (t *TencentCOS) MultipartUploadFile(file *multipart.FileHeader, keyPrefix ...string) (string, string, string, error) {
	client := NewClient(t.cfg)
	f, openError := file.Open()
	if openError != nil {
		log.Zaplog().Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKeyBuild := make([]string, 0)
	if t.cfg.PrefixPath != "" {
		fileKeyBuild = append(fileKeyBuild, t.cfg.PrefixPath)
	}
	if len(keyPrefix) > 0 {
		fileKeyBuild = append(fileKeyBuild, strings.Join(keyPrefix, "/"))
	}
	fileKeyBuild = append(fileKeyBuild, fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename))
	fileKey := strings.Join(fileKeyBuild, "/")

	md5 := fileMd5(f)
	_, err := client.Object.Put(context.Background(), fileKey, f, nil)
	if err != nil {
		panic(err)
	}

	if t.cfg.CdnURL != "" {
		return t.cfg.CdnURL + "/" + fileKey, fileKey, md5, nil
	}
	return client.BaseURL.BucketURL.Host + "/" + fileKey, fileKey, md5, nil
}

func (t *TencentCOS) UploadFile(file *os.File, keyPrefix ...string) (reqPath, fileKey, md5 string, err error) {
	panic("未实现上传方法")
}

func (t *TencentCOS) DeleteFile(key string) error {
	client := NewClient(t.cfg)
	_, err := client.Object.Delete(context.Background(), key)
	if err != nil {
		log.Zaplog().Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

func NewClient(cfg config.File) *cos.Client {
	urlStr, _ := url.Parse("https://" + cfg.Bucket + ".cos." + cfg.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.ID,
			SecretKey: cfg.Key,
		},
	})
	return client
}
