package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"

	"go.uber.org/zap"
)

type Qiniu struct {
	cfg config.File
}

func (q *Qiniu) MultipartUploadFile(file *multipart.FileHeader, keyPrefix ...string) (string, string, string, error) {
	putPolicy := storage.PutPolicy{Scope: q.cfg.Bucket}
	mac := qbox.NewMac(q.cfg.ID, q.cfg.Key)
	upToken := putPolicy.UploadToken(mac)
	cfg := q.qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		log.Zaplog().Error("function file.Open() Filed", zap.Any("err", openError.Error()))

		return "", "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	pathKeyBuild := []string{}
	if q.cfg.PrefixPath != "" {
		pathKeyBuild = append(pathKeyBuild, q.cfg.PrefixPath)
	}
	if len(keyPrefix) > 0 {
		pathKeyBuild = append(pathKeyBuild, keyPrefix...)
	}
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	pathKeyBuild = append(pathKeyBuild, fileName)
	fileKey := strings.Join(pathKeyBuild, "/")
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		log.Zaplog().Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}

	md5 := fileMd5(f)
	return q.cfg.CdnURL + ret.Key, fileName, md5, nil
}

func (q *Qiniu) UploadFile(file *os.File, keyPrefix ...string) (reqPath, fileKey, md5 string, err error) {

	putPolicy := storage.PutPolicy{Scope: q.cfg.Bucket}
	mac := qbox.NewMac(q.cfg.ID, q.cfg.Key)
	upToken := putPolicy.UploadToken(mac)
	cfg := q.qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	pathKeyBuild := []string{}
	if q.cfg.PrefixPath != "" {
		pathKeyBuild = append(pathKeyBuild, q.cfg.PrefixPath)
	}
	if len(keyPrefix) > 0 {
		pathKeyBuild = append(pathKeyBuild, keyPrefix...)
	}
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Name())
	pathKeyBuild = append(pathKeyBuild, fileName)

	filePathName := strings.Join(pathKeyBuild, "/")

	stat, _ := file.Stat()
	putErr := formUploader.Put(context.Background(), &ret, upToken, filePathName, file, stat.Size(), &putExtra)
	if putErr != nil {
		log.Zaplog().Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}

	md5 = fileMd5(file)
	return q.cfg.CdnURL + ret.Key, fileName, md5, nil
}

func (q *Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(q.cfg.ID, q.cfg.Key)
	cfg := q.qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(q.cfg.Bucket, key); err != nil {
		log.Zaplog().Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

func (q *Qiniu) UploadTokenGet(key string) (token string, fileKey string, err error) {
	mac := credentials.NewCredentials(q.cfg.ID, q.cfg.Key)
	putPolicy, err := uptoken.NewPutPolicy(q.cfg.Bucket, time.Now().Add(1*time.Hour))
	if err != nil {
		return "", "", err
	}

	pathKeyBuild := []string{}
	if q.cfg.PrefixPath != "" {
		pathKeyBuild = append(pathKeyBuild, q.cfg.PrefixPath)
	}
	pathKeyBuild = append(pathKeyBuild, key)
	fileKey = "/" + strings.Join(pathKeyBuild, "/")
	token, err = uptoken.NewSigner(putPolicy, mac).GetUpToken(context.Background())
	return token, strings.Join(pathKeyBuild, "/"), err
}

func (q *Qiniu) qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      q.cfg.CdnURL != "", // TODO 这里似乎不匹配
		UseCdnDomains: q.cfg.CdnURL != "", // TODO 这里似乎不匹配
	}
	switch q.cfg.Region {
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
