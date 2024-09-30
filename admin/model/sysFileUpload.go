package model

import (
	"errors"

	"github.com/mangk/adminBox/db"
	"gorm.io/gorm"
)

type SysFileUpload struct {
	Model
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号
}

func (SysFileUpload) TableName() string {
	return "sys_file_upload"
}

func (s SysFileUpload) Upload(file *SysFileUpload) error {
	return db.DB().Create(file).Error
}

func (s SysFileUpload) FindFile(id int) (SysFileUpload, error) {
	var file SysFileUpload
	err := db.DB().Where("id = ?", id).First(&file).Error
	return file, err
}

func (s SysFileUpload) EditFileName(file SysFileUpload) (err error) {
	var fileFromDb SysFileUpload
	return db.DB().Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}

func (s SysFileUpload) DeleteFile(file SysFileUpload) (err error) {
	// var fileFromDb SysFileUpload
	// fileFromDb, err = s.FindFile(file.ID)
	// if err != nil {
	// 	return
	// }

	if err := db.DB().Where("id = ?", file.ID).Delete(&file).Error; err != nil {
		return err
	}

	// oss := upload.NewOss()
	// if err = oss.DeleteFile(fileFromDb.Key); err != nil {
	// 	return errors.New("文件删除失败")
	// }
	return err
}

func (s SysFileUpload) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (file ExaFile, err error) {
	var cfile ExaFile
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal

	if errors.Is(db.DB().Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&file).Error, gorm.ErrRecordNotFound) {
		err = db.DB().Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&file, cfile).Error
		return file, err
	}
	cfile.IsFinish = true
	cfile.FilePath = file.FilePath
	err = db.DB().Create(&cfile).Error
	return cfile, err
}

// file struct, 文件结构体
type ExaFile struct {
	Model
	FileName   string
	FileMd5    string
	FilePath   string
	ChunkTotal int
	IsFinish   bool
}
