package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
	"go.uber.org/zap"
)

type AliyunOSS struct {
	cfg config.File
}

func (a *AliyunOSS) MultipartUploadFile(file *multipart.FileHeader, keyPrefix ...string) (string, string, string, int64, error) {
	bucket, err := NewBucket(a.cfg)
	c := bucket.GetConfig()
	fmt.Println(c)
	if err != nil {
		log.Zaplog().Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", "", file.Size, errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		log.Zaplog().Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", "", file.Size, errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKeyBuild := make([]string, file.Size)
	if a.cfg.PrefixPath != "" {
		fileKeyBuild = append(fileKeyBuild, a.cfg.PrefixPath)
	}
	fileKeyBuild = append(fileKeyBuild, fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename))
	fileKey := strings.Join(fileKeyBuild, "/")

	md5 := fileMd5(f)
	// 上传文件流。
	err = bucket.PutObject(fileKey, f)
	if err != nil {
		log.Zaplog().Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", "", file.Size, errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	if a.cfg.CdnURL != "" {
		return a.cfg.CdnURL + "/" + fileKey, fileKey, md5, file.Size, nil
	}
	return bucket.BucketName + "." + bucket.Client.Config.Endpoint + "/" + fileKey, fileKey, md5, file.Size, nil
}

func (a *AliyunOSS) UploadFile(file *os.File, keyPrefix ...string) (reqPath, fileKey, md5 string, err error) {
	panic("未实现上传方法")
}

func (a *AliyunOSS) DeleteFile(key string) error {
	bucket, err := NewBucket(a.cfg)
	if err != nil {
		log.Zaplog().Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		log.Zaplog().Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

func (a *AliyunOSS) UploadTokenGet(key string, uuid string) (token string, fileKey string, err error) {
	panic("未实现上传方法")
}

func NewBucket(cfg config.File) (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(cfg.Region, cfg.ID, cfg.Key)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(cfg.Bucket)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
