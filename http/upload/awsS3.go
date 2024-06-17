package upload

import (
	"errors"
	"fmt"
	"github.com/mangk/adminX/config"
	"github.com/mangk/adminX/log"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

type AwsS3 struct {
	cfg config.File
}

//@author: [WqyJh](https://github.com/WqyJh)
//@object: *AwsS3
//@function: UploadFile
//@description: Upload file to Aws S3 using aws-sdk-go. See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html#s3-examples-bucket-ops-upload-file-to-bucket
//@param: file *multipart.FileHeader
//@return: string, string, error

func (a *AwsS3) UploadFile(file *multipart.FileHeader) (string, string, error) {
	session := a.newSession()
	uploader := s3manager.NewUploader(session)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := a.cfg.PrefixPath + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		log.Zaplog().Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(a.cfg.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		log.Zaplog().Error("function uploader.Upload() Filed", zap.Any("err", err.Error()))
		return "", "", err
	}

	return a.cfg.PrefixPath + "/" + filename, fileKey, nil
}

//@author: [WqyJh](https://github.com/WqyJh)
//@object: *AwsS3
//@function: DeleteFile
//@description: Delete file from Aws S3 using aws-sdk-go. See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html#s3-examples-bucket-ops-delete-bucket-item
//@param: file *multipart.FileHeader
//@return: string, string, error

func (a *AwsS3) DeleteFile(key string) error {
	session := a.newSession()
	svc := s3.New(session)
	filename := a.cfg.PrefixPath + "/" + key
	bucket := a.cfg.Bucket

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Zaplog().Error("function svc.DeleteObject() Filed", zap.Any("err", err.Error()))
		return errors.New("function svc.DeleteObject() Filed, err:" + err.Error())
	}

	_ = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return nil
}

// newSession Create S3 session
func (a *AwsS3) newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String(a.cfg.Region),
		Endpoint: aws.String(a.cfg.Bucket), //minio在这里设置地址,可以兼容
		//S3ForcePathStyle: aws.Bool(core.Config().AwsS3.S3ForcePathStyle), // TODO 这里似乎不匹配
		//DisableSSL:       aws.Bool(core.Config().AwsS3.DisableSSL), // TODO 这里似乎不匹配
		Credentials: credentials.NewStaticCredentials(
			a.cfg.ID,
			a.cfg.Key,
			"",
		),
	})
	return sess
}
