package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/admin/model"
	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/http/request"
	"github.com/mangk/adminBox/http/response"
)

func Role(ctx *gin.Context) {
	req := request.PublicRequest(ctx)

	var count int64
	list := []model.SysRole{}

	query := db.DB().Model(list) // TODO 补充搜索条件

	if err := query.Count(&count).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if count > 0 {
		if err := query.Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Find(&list).Error; err != nil {
			response.FailWithError(ctx, err)
			return
		}
	}

	response.OkWithPageData(ctx, count, list)
}

func RoleCreate(ctx *gin.Context) {
	req := model.SysRole{}
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

func RoleEdit(ctx *gin.Context) {
	req := model.SysRole{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	update := make(map[string]interface{})
	update["name"] = req.Name
	update["description"] = req.Description

	if err := db.DB().Model(&req).Where("id = ?", req.ID).Updates(update).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.Ok(ctx)
}

func RoleDetail(ctx *gin.Context) {
	req := model.SysRole{}
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

func RoleDelete(ctx *gin.Context) {
	req := model.SysRole{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if err := db.DB().Delete(&req).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.Ok(ctx)
}

func RoleAll(ctx *gin.Context) {
	list := []model.SysRole{}
	if err := db.DB().Find(&list).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.OkWithData(ctx, list)
}
