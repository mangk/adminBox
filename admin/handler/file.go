package handler

import (
	"mime/multipart"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/admin/model"
	"github.com/mangk/adminX/config"
	"github.com/mangk/adminX/db"
	"github.com/mangk/adminX/http/request"
	"github.com/mangk/adminX/http/response"
	"github.com/mangk/adminX/http/upload"
	"github.com/mangk/adminX/log"
)

func FileGetUploadLimit(ctx *gin.Context) {
	uploadCfg := config.FileCfg()
	resp := gin.H{}
	for name, cfg := range uploadCfg {
		resp[name] = gin.H{"name": cfg.Name, "limit": cfg.Limit * 1024, "driver": name}
	}
	response.OkWithData(ctx, resp)
}

func FileUpload(ctx *gin.Context) {
	var file model.SysFileUpload
	noSave := ctx.DefaultQuery("noSave", "0")
	driver := ctx.DefaultQuery("driver", "default")
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Errorf("接收文件失败! %s", err)
		response.FailWithMsg(ctx, "接收文件失败")
		return
	}
	file, err = uploadFile(header, noSave, driver, request.JWTUserId(ctx)) // 文件上传后拿到文件路径
	if err != nil {
		log.Errorf("修改数据库链接失败! %s", err)
		response.FailWithMsg(ctx, "修改数据库链接失败")
		return
	}
	response.OkWithDetail(ctx, "上传成功", file)
}

func FileList(ctx *gin.Context) {
	req := request.PublicRequest(ctx)

	var count int64
	list := []model.SysFileUpload{}

	query := db.DB().Model(list).Where("cb = ?", request.JWTUserId(ctx))

	if err := query.Count(&count).Error; err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	if count > 0 {
		if err := query.Order("id desc").Limit(req.PageSize).Offset(req.PageSize * (req.Page - 1)).Find(&list).Error; err != nil {
			response.FailWithMsg(ctx, err.Error())
			return
		}
	}

	for i := range list {
		list[i].Url = "http://127.0.0.1:8910" + list[i].Url // TODO 这里结合config 处理
	}

	response.OkWithPageData(ctx, count, list)
}

func FileDelete(c *gin.Context) {
	var file model.SysFileUpload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMsg(c, err.Error())
		return
	}
	if err := file.DeleteFile(file); err != nil {
		response.FailWithMsg(c, "删除失败")
		return
	}
	response.OkWithMsg(c, "删除成功")
}

func FileEdit(ctx *gin.Context) {
	var file model.SysFileUpload
	err := ctx.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}
	err = file.EditFileName(file)
	if err != nil {
		response.FailWithMsg(ctx, "编辑失败")
		return
	}
	response.OkWithMsg(ctx, "编辑成功")
}

func uploadFile(header *multipart.FileHeader, noSave, driver string, cb int) (file model.SysFileUpload, err error) {
	oss := upload.NewOss(driver)
	filePath, key, _, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		t := model.LocalTime(time.Now())
		s := strings.Split(header.Filename, ".")
		f := model.SysFileUpload{
			Model: model.Model{Cb: cb, Ct: &t},
			Url:   filePath,
			Name:  header.Filename,
			Tag:   s[len(s)-1],
			Key:   key,
		}
		return f, (model.SysFileUpload{}).Upload(&f)
	}
	return
}
