package handler

import (
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mangk/adminBox/admin/model"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/http/request"
	"github.com/mangk/adminBox/http/response"
	"github.com/mangk/adminBox/http/upload"
	"github.com/mangk/adminBox/log"
)

func FileGroupTree(ctx *gin.Context) {
	tree, err := model.SysFileGroup{}.Tree(request.JWTLoginUserId(ctx))
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.OkWithData(ctx, []gin.H{{"name": "默认分组", "id": 0, "children": tree}})
}

func FileGroupCreate(ctx *gin.Context) {
	req := model.SysFileGroup{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}
	req.Cb = request.JWTLoginUserId(ctx)
	db.DB().Create(&req)
	response.OkWithMsg(ctx, "创建成功")
}

func FileGroupDelete(ctx *gin.Context) {
	req := model.SysFileGroup{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}
	if req.ID == 0 {
		response.FailWithMsg(ctx, "参数错误")
		return
	}

	// 检查是否有子分组
	var count int64
	db.DB().Model(model.SysFileGroup{}).Where("parent_id =?", req.ID).Count(&count)
	if count > 0 {
		response.FailWithMsg(ctx, "分组下存在子分组，无法删除")
		return
	}
	// 检查是否有文件
	var count2 int64
	db.DB().Model(model.SysFile{}).Where("group_id =?", req.ID).Count(&count2)
	if count2 > 0 {
		response.FailWithMsg(ctx, "分组下存在文件，无法删除")
		return
	}

	if err := db.DB().Delete(&req).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.OkWithMsg(ctx, "删除成功")
}

func FileGetUploadLimit(ctx *gin.Context) {
	uploadCfg := config.FileCfg()
	resp := gin.H{}
	for name, cfg := range uploadCfg {
		resp[name] = gin.H{"name": cfg.Name, "limit": cfg.Limit * 1024, "driver": name, "cdn": cfg.CdnURL, "prefix": cfg.PrefixPath}
	}
	response.OkWithData(ctx, resp)
}

func FileUpload(ctx *gin.Context) {
	var file model.SysFile
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

func FileUploadToken(ctx *gin.Context) {
	// TODO 这里增加储存桶，避免 token 过度下发

	req := struct {
		FileName string `json:"file_name"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, "参数错误")
		return
	}

	userId := request.JWTLoginUserId(ctx)
	driver := ctx.DefaultQuery("driver", "default")
	cfg := config.FileCfg()[driver]
	oss := upload.NewOss(driver)
	uuid := uuid.NewString()
	token, key, err := oss.UploadTokenGet(fmt.Sprintf("%d/%d_%s", userId, time.Now().Unix(), req.FileName), uuid)
	if err != nil {
		response.FailWithMsg(ctx, "获取上传凭证失败")
		return
	}

	// s := strings.Split(req.FileName, ".")
	// file := model.SysFile{
	// 	Model: model.Model{Cb: userId},
	// 	Name:  req.FileName,
	// 	Tag:   s[len(s)-1],
	// 	Key:   key,
	// 	UUID:  uuid,
	// }
	// if err := db.DB().Omit("group_id").Create(&file).Error; err != nil {
	// 	response.FailWithMsg(ctx, "保存文件信息失败")
	// 	return
	// }
	response.OkWithDetail(ctx, "获取上传凭证成功", gin.H{"token": token, "key": key, "driver": cfg.Driver})
}

func FileSaveUploadFileInfo(ctx *gin.Context) {
	req := struct {
		Path   string `json:"path"`
		Driver string `json:"driver"`
		Group  int    `json:"group"`
		Key    string `json:"key"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, "参数错误")
		return
	}

	userId := request.JWTLoginUserId(ctx)
	pathInfo := strings.Split(req.Path, "/")
	lenPathInfo := len(pathInfo)
	if lenPathInfo > 1 {
		tagInfo := strings.Split(pathInfo[lenPathInfo-1], ".")

		uuid := uuid.NewString()
		file := model.SysFile{
			Model:   model.Model{Cb: userId},
			Name:    pathInfo[lenPathInfo-1],
			Tag:     tagInfo[len(tagInfo)-1],
			Key:     req.Key,
			UUID:    uuid,
			Url:     req.Path,
			GroupId: req.Group,
		}
		if err := db.DB().Create(&file).Error; err != nil {
			response.FailWithMsg(ctx, "保存文件信息失败")
			return
		}
	}
}

func FileUploadCallback(ctx *gin.Context) {
	req := struct {
		Key    string `json:"key"`
		Hash   string `json:"hash"`
		Fsize  int64  `json:"fsize"`
		Bucket string `json:"bucket"`
		Name   string `json:"name"`
		UUID   string `json:"uuid"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, "参数错误")
		return
	}
	file := model.SysFile{}
	if err := db.DB().Where("uuid =?", req.UUID).First(&file).Error; err != nil {
		response.FailWithMsg(ctx, "文件不存在")
		return
	}
	if file.Key != req.Key {
		response.FailWithMsg(ctx, "文件上传失败")
		return
	}
	file.Url = fmt.Sprintf("%s/%s", strings.TrimRight(config.FileCfg()["default"].CdnURL, "/"), strings.TrimLeft(file.Key, "/"))
	file.Size = req.Fsize
	db.DB().Save(&file)
	response.OkWithMsg(ctx, "文件上传成功")
}

func FileList(ctx *gin.Context) {
	req := request.PublicRequest(ctx)

	var count int64
	list := []model.SysFile{}

	query := db.DB().Model(list).Where("cb = ?", request.JWTLoginUserId(ctx))
	if qt, has := req.Query["tag"]; has && len(qt.([]interface{})) > 0 {
		query = query.Where("tag in ?", qt)
	}
	if qg, has := req.Query["group_id"]; has {
		query = query.Where("group_id = ?", qg)
	}

	if name, has := req.Query["name"]; has {
		query = query.Where("name like ?", "%"+name.(string)+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if count > 0 {
		if err := query.Order("id desc").Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Preload("GroupInfo").Find(&list).Error; err != nil {
			response.FailWithError(ctx, err)
			return
		}
	}

	host := strings.TrimRight(config.ServerCfg().RunAt, "/")
	style := config.FileCfg()["default"].Style
	for i := range list {
		if strings.HasPrefix(list[i].Url, "http") {
			continue
		}
		list[i].Url = host + list[i].Url // TODO 这里结合config 处理
		if style != "" {
			list[i].Url += style
		}
	}

	response.OkWithPageData(ctx, count, list)
}

func FileDelete(c *gin.Context) {
	var file model.SysFile
	err := c.ShouldBindBodyWithJSON(&file)
	if err != nil {
		req := struct {
			Id []int `json:"id"`
		}{}
		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			response.FailWithMsg(c, "参数错误")
			return
		}
		if err := db.DB().Where("cb =?", request.JWTLoginUserId(c)).Where("id in?", req.Id).Delete(&file).Error; err != nil {
			response.FailWithMsg(c, "删除失败")
			return
		}
	} else {
		if err := file.DeleteFile(file); err != nil {
			response.FailWithMsg(c, "删除失败")
			return
		}
	}
	response.OkWithMsg(c, "删除成功")
}

func FileMove(ctx *gin.Context) {
	req := struct {
		Ids     []int `json:"ids"`
		GroupId int   `json:"group_id"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	group := model.SysFileGroup{}
	db.DB().Where("cb = ?", request.JWTLoginUserId(ctx)).First(&group, req.GroupId)
	if group.ID == 0 {
		if req.GroupId == 0 {
			db.DB().Debug().Model(&model.SysFile{}).Where("cb =?", request.JWTLoginUserId(ctx)).Where("id in ?", req.Ids).Update("group_id", nil)
			response.OkWithMsg(ctx, "移动成功")
			return
		}
		response.FailWithMsg(ctx, "分组不存在")
		return
	}
	db.DB().Debug().Model(&model.SysFile{}).Where("cb =?", request.JWTLoginUserId(ctx)).Where("id in ?", req.Ids).Update("group_id", req.GroupId)
	response.OkWithMsg(ctx, "移动成功")
}

func FileEdit(ctx *gin.Context) {
	var file model.SysFile
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

func uploadFile(header *multipart.FileHeader, noSave, driver string, cb int) (file model.SysFile, err error) {
	oss := upload.NewOss(driver)
	filePath, key, _, fsize, uploadErr := oss.MultipartUploadFile(header, fmt.Sprintf("%d", cb))
	if uploadErr != nil {
		panic(uploadErr)
	}
	if noSave == "0" {
		t := model.LocalTime(time.Now())
		s := strings.Split(header.Filename, ".")
		f := model.SysFile{
			Model: model.Model{Cb: cb, Ct: &t},
			Url:   filePath,
			Name:  header.Filename,
			Tag:   s[len(s)-1],
			Key:   key,
			UUID:  uuid.NewString(),
			Size:  fsize,
		}
		return f, (model.SysFile{}).Upload(&f)
	}
	return
}
