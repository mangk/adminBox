package model

import (
	"errors"

	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/db"
	"gorm.io/gorm"
)

type SysFileGroup struct {
	Model
	ParentId int            `json:"parent_id" gorm:"comment:父级id"` // 父级id
	Name     string         `json:"name" gorm:"comment:路径"`        // 路径
	Desc     string         `json:"desc" gorm:"comment:描述"`        // 描述
	Files    []SysFile      `json:"files" gorm:"foreignKey:GroupId;references:ID"`
	Children []SysFileGroup `json:"children,omitempty" gorm:"-"`
}

func (SysFileGroup) TableName() string {
	return config.DBCfg()["default"].Prefix + "sys_file_group"
}

func (s SysFileGroup) Tree(userId int) ([]SysFileGroup, error) {
	var list []SysFileGroup
	err := db.DB().Where("cb =?", userId).Find(&list).Error
	if err != nil {
		return nil, err
	}

	treeMap := make(map[int][]SysFileGroup)
	for _, v := range list {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	resMenuList := treeMap[0]
	for i := 0; i < len(resMenuList); i++ {
		err = s.getBaseChildrenList(&resMenuList[i], treeMap)
	}
	return resMenuList, err
}

func (s SysFileGroup) getBaseChildrenList(menu *SysFileGroup, treeMap map[int][]SysFileGroup) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = s.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

type SysFile struct {
	Model
	GroupId   int          `json:"group_id" gorm:"comment:路径id"` // 路径id
	GroupInfo SysFileGroup `json:"group_info" gorm:"foreignKey:GroupId;references:ID"`
	Name      string       `json:"name" gorm:"comment:文件名"`    // 文件名
	Url       string       `json:"url" gorm:"comment:文件地址"`    // 文件地址
	Tag       string       `json:"tag" gorm:"comment:文件标签"`    // 文件标签
	Key       string       `json:"key" gorm:"comment:编号"`      // 编号
	UUID      string       `json:"uuid" gorm:"comment:回调uuid"` // 回调uuid
	Size      int64        `json:"size" gorm:"comment:文件大小"`   // 文件大小
}

func (SysFile) TableName() string {
	return config.DBCfg()["default"].Prefix + "sys_file"
}

func (s SysFile) Upload(file *SysFile) error {
	return db.DB().Omit("group_id").Create(file).Error
}

func (s SysFile) FindFile(id int) (SysFile, error) {
	var file SysFile
	err := db.DB().Where("id = ?", id).First(&file).Error
	return file, err
}

func (s SysFile) EditFileName(file SysFile) (err error) {
	var fileFromDb SysFile
	return db.DB().Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}

func (s SysFile) DeleteFile(file SysFile) (err error) {
	// var fileFromDb SysFileUpload
	// fileFromDb, err = s.FindFile(file.ID)
	// if err != nil {
	// 	return
	// }

	if err = db.DB().Where("id = ?", file.ID).Delete(&file).Error; err != nil {
		return err
	}

	// oss := upload.NewOss()
	// if err = oss.DeleteFile(fileFromDb.Key); err != nil {
	// 	return errors.New("文件删除失败")
	// }
	return err
}

func (s SysFile) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (file ExaFile, err error) {
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
