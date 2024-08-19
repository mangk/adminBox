package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/admin/model"
	"github.com/mangk/adminX/db"
	"github.com/mangk/adminX/http/response"
)

func Menu(ctx *gin.Context) {
	req := struct {
		LoadSystem bool `json:"loadSystem"`
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if tree, err := (model.SysMenu{}).Tree(req.LoadSystem, false, true); err == nil {
		response.OkWithData(ctx, tree)
	} else {
		response.FailWithError(ctx, err)
	}
}

func MenuCreate(ctx *gin.Context) {
	req := model.SysMenu{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if err := db.DB().Create(&req).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.OkWithData(ctx, req.ID)
}

func MenuDetail(ctx *gin.Context) {
	req := model.SysMenu{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if err := db.DB().Where("id = ?", req.ID).First(&req).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.OkWithData(ctx, req)
}

func MenuEdit(ctx *gin.Context) {
	req := model.SysMenu{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	update := make(map[string]interface{})
	//update["ub"] = req.
	update["pid"] = req.Pid
	update["name"] = req.Name
	update["path"] = req.Path
	update["hidden"] = req.Hidden
	update["component"] = req.Component
	update["sort"] = req.Sort
	update["title"] = req.Title
	update["keep_alive"] = req.KeepAlive
	update["default_menu"] = req.DefaultMenu
	update["icon"] = req.Icon
	update["auto_close"] = req.AutoClose
	update["sc_path"] = req.SCPath
	al, _ := json.Marshal(req.ActionList)
	update["action_list"] = string(al)

	if err := db.DB().Model(&req).Where("id = ?", req.ID).Updates(update).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.Ok(ctx)
}
func MenuDelete(ctx *gin.Context) {
	req := model.SysMenu{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	var hasCount int64
	db.DB().Model(req).Where("id = ? or pid = ?", req.ID, req.ID).Count(&hasCount)

	if hasCount > 1 {
		response.FailWithMsg(ctx, "当前菜单存在子菜单，无法删除")
		return
	}
	if hasCount <= 0 {
		response.FailWithMsg(ctx, "数据不存在")
		return
	}

	if err := db.DB().Delete(&req).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.Ok(ctx)
}
