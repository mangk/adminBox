package handler

import (
	"mime/multipart"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/admin/model"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/http/request"
	"github.com/mangk/adminBox/http/response"
	"github.com/mangk/adminBox/http/upload"
	"github.com/mangk/adminBox/log"
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
	file, err = uploadFile(header, noSave, driver, request.JWTLoginUserId(ctx)) // 文件上传后拿到文件路径
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

	query := db.DB().Model(list).Where("cb = ?", request.JWTLoginUserId(ctx))
	if qt, has := req.Query["tag"]; has && len(qt.([]interface{})) > 0 {
		query = query.Where("tag in ?", qt)
	}

	if err := query.Count(&count).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if count > 0 {
		if err := query.Order("id desc").Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Find(&list).Error; err != nil {
			response.FailWithError(ctx, err)
			return
		}
	}

	host := strings.TrimRight(config.ServerCfg().RunAt, "/")
	for i := range list {
		if strings.HasPrefix(list[i].Url, "http") {
			continue
		}
		list[i].Url = host + list[i].Url // TODO 这里结合config 处理
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
		response.FailWithError(ctx, err)
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
	filePath, key, _, uploadErr := oss.MultipartUploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
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
