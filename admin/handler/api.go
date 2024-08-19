package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/admin/model"
	"github.com/mangk/adminX/db"
	"github.com/mangk/adminX/http/request"
	"github.com/mangk/adminX/http/response"
)

func Api(ctx *gin.Context) {
	req := request.PublicRequest(ctx)

	var count int64
	list := []model.SysApi{}

	query := db.DB().Model(list) // TODO 补充搜索条件

	if err := query.Count(&count).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if count > 0 {
		if err := query.Limit(req.PageSize).Offset(req.PageSize * (req.Page - 1)).Find(&list).Error; err != nil {
			response.FailWithError(ctx, err)
			return
		}
	}

	tranMap := model.SysMenu{}.TranMap()
	for i, api := range list {
		if menuName, ok := tranMap[api.MenuId]; ok {
			list[i].MenuName = menuName
		}
	}

	response.OkWithPageData(ctx, count, list)
}

func ApiCreate(ctx *gin.Context) {
	req := model.SysApi{}
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

func ApiEdit(ctx *gin.Context) {
	req := model.SysApi{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	update := make(map[string]interface{})
	//update["ub"] = req.
	update["menu_id"] = req.MenuId
	update["name"] = req.Name
	update["description"] = req.Description
	update["path"] = req.Path
	update["method"] = req.Method

	if err := db.DB().Model(&req).Where("id = ?", req.ID).Updates(update).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.Ok(ctx)
}

func ApiDetail(ctx *gin.Context) {
	req := model.SysApi{}
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

func ApiDelete(ctx *gin.Context) {
	req := model.SysApi{}
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
