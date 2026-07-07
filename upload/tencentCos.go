package upload

import (
	"context"
	"encoding/json"
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
	"github.com/tidwall/sjson"

	"github.com/tencentyun/cos-go-sdk-v5"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"go.uber.org/zap"
)

type TencentCOS struct {
	cfg config.File
}

func (t *TencentCOS) MultipartUploadFile(file *multipart.FileHeader, keyPrefix ...string) (string, string, string, int64, error) {
	client := NewClient(t.cfg)
	f, openError := file.Open()
	if openError != nil {
		log.Zaplog().Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", "", file.Size, errors.New("function file.Open() Filed, err:" + openError.Error())
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
		return t.cfg.CdnURL + "/" + fileKey, fileKey, md5, file.Size, nil
	}
	return client.BaseURL.BucketURL.Host + "/" + fileKey, fileKey, md5, file.Size, nil
}

func (t *TencentCOS) UploadFile(file *os.File, keyPrefix ...string) (reqPath, fileKey, md5 string, err error) {
	panic("未实现上传方法")
}

func (t *TencentCOS) UploadTokenGet(key string, uuid string) (token string, fileKey string, err error) {

	pathKeyBuild := []string{}
	if t.cfg.PrefixPath != "" {
		pathKeyBuild = append(pathKeyBuild, t.cfg.PrefixPath)
	}
	pathKeyBuild = append(pathKeyBuild, key)
	fileKey = strings.Join(pathKeyBuild, "/")

	client := sts.NewClient(t.cfg.ID, t.cfg.Key, nil)

	// TODO 图片类型
	// TODO 限制大小
	// TODO 确定这里的内容是否准确
	opt := &sts.CredentialOptions{
		DurationSeconds: 1800,
		Region:          t.cfg.Region,
		Policy: &sts.CredentialPolicy{
			Version: "2.0",
			Statement: []sts.CredentialPolicyStatement{{
				Action: []string{
					// 简单上传
					"name/cos:PutObject",

					// 简单上传
					"name/cos:InitiateMultipartUpload",
					"name/cos:ListMultipartUploads",
					"name/cos:ListParts",
					"name/cos:UploadPart",
					"name/cos:CompleteMultipartUpload",
				},
				Effect: "allow",
				Resource: []string{
					// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
					// 存储桶的命名格式为 BucketName-APPID，此处填写的 bucket 必须为此格式
					"qcs::cos:" + t.cfg.Region + ":uid/" + strings.Split(t.cfg.Bucket, "-")[1] + ":" + t.cfg.Bucket + "/" + fileKey,
				},
				Condition: map[string]map[string]interface{}{
					"string_like_if_exist": {
						// 只允许上传 content-type 为图片类型
						"cos:content-type": "image/*", // TODO 这里类型应该有视频和图片和其他的
					},
					"numeric_less_than_equal": {
						// 上传大小限制不能超过5MB(只对简单上传生效)
						"cos:content-length": 50 * 1024 * 1024,
					},
				},
			}},
		},
	}

	res, err := client.GetCredential(opt)
	if err != nil {
		return "", "", err
	}

	bytesRes, _ := json.Marshal(res)
	bytesRes, _ = sjson.SetBytes(bytesRes, "bucket", t.cfg.Bucket)
	bytesRes, _ = sjson.SetBytes(bytesRes, "region", t.cfg.Region)
	bytesRes, err = sjson.SetBytes(bytesRes, "key", fileKey)
	return string(bytesRes), fileKey, err
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
