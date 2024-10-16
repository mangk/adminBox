package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mangk/adminBox/admin/model"
	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/http/request"
	"github.com/mangk/adminBox/http/response"
)

func User(ctx *gin.Context) {
	req := request.PublicRequest(ctx)

	var count int64
	list := []model.SysUser{}

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

func UserCreate(ctx *gin.Context) {
	req := model.SysUser{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}
	req.UUID = uuid.New()

	if err := db.DB().Create(&req).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.OkWithData(ctx, req.ID)
}

func UserEdit(ctx *gin.Context) {
	req := model.SysUser{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if err := (model.SysUser{}).Update(req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if err := (model.SysCasbinRole{}).UpdateCasbin(req.ID); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.Ok(ctx)
}

func UserChangePassord(ctx *gin.Context) {

}

func UserDetail(ctx *gin.Context) {
	req := model.SysUser{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	user, err := model.SysUser{}.Detail(req.ID)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.OkWithData(ctx, user)
}

func UserDelete(ctx *gin.Context) {
	req := model.SysUser{}
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
