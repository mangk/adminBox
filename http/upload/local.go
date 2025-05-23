package upload

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"
	"github.com/mangk/adminBox/util"

	"go.uber.org/zap"
)

type Local struct {
	cfg config.File
}

func (l *Local) MultipartUploadFile(file *multipart.FileHeader, keyPrefix ...string) (string, string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = util.MD5V([]byte(name))
	// 拼接新文件名
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	if len(keyPrefix) > 0 {
		filename = strings.Join(keyPrefix, "_") + "_" + filename
	}
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(l.cfg.StorePath, os.ModePerm)
	if mkdirErr != nil {
		log.Zaplog().Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := l.cfg.StorePath + "/" + filename

	fileKeyBuild := make([]string, 0)
	if l.cfg.PrefixPath != "" {
		fileKeyBuild = append(fileKeyBuild, l.cfg.PrefixPath)
	}
	fileKeyBuild = append(fileKeyBuild, filename)
	fileKey := strings.Join(fileKeyBuild, "/")

	f, openError := file.Open() // 读取文件
	if openError != nil {
		log.Zaplog().Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		log.Zaplog().Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	md5 := fileMd5(f)
	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		log.Zaplog().Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}

	if l.cfg.CdnURL != "" {
		return l.cfg.CdnURL + "/" + fileKey, filename, md5, nil
	}
	return "/" + fileKey, filename, md5, nil
}

func (l *Local) UploadFile(file *os.File, keyPrefix ...string) (reqPath, fileKey, md5 string, err error) {
	panic("未实现上传方法")
}

func (l *Local) DeleteFile(key string) error {
	p := l.cfg.StorePath + "/" + key
	if strings.Contains(p, l.cfg.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
