package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/admin/model"
	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/http/response"
)

func Department(ctx *gin.Context) {
	if tree, err := (model.SysDepartment{}).All(); err == nil {
		response.OkWithData(ctx, tree)
	} else {
		response.FailWithError(ctx, err)
	}
}

func DepartmentCreate(ctx *gin.Context) {
	req := model.SysDepartment{}
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

func DepartmentDetail(ctx *gin.Context) {
	req := model.SysDepartment{}
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

func DepartmentEdit(ctx *gin.Context) {
	req := model.SysDepartment{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	update := make(map[string]interface{})
	//update["ub"] = req.
	update["pid"] = req.Pid
	update["name"] = req.Name
	update["description"] = req.Description

	if err := db.DB().Model(&req).Where("id = ?", req.ID).Updates(update).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.Ok(ctx)
}
func DepartmentDelete(ctx *gin.Context) {
	req := model.SysDepartment{}
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
