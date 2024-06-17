package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mangk/adminX/admin/model"
	"github.com/mangk/adminX/db"
	"github.com/mangk/adminX/http/request"
	"github.com/mangk/adminX/http/response"
)

func User(ctx *gin.Context) {
	req := request.PublicRequest(ctx)

	var count int64
	list := []model.SysUser{}

	query := db.DB().Model(list) // TODO 补充搜索条件

	if err := query.Count(&count).Error; err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	if count > 0 {
		if err := query.Limit(req.PageSize).Offset(req.PageSize * (req.Page - 1)).Find(&list).Error; err != nil {
			response.FailWithMsg(ctx, err.Error())
			return
		}
	}

	response.OkWithPageData(ctx, count, list)
}

func UserCreate(ctx *gin.Context) {
	req := model.SysUser{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}
	req.UUID = uuid.New()

	if err := db.DB().Create(&req).Error; err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, req.ID)
}

func UserEdit(ctx *gin.Context) {
	req := model.SysUser{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	if err := (model.SysUser{}).Update(req); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	if err := (model.SysCasbinRole{}).UpdateCasbin(req.ID); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}

func UserChangePassord(ctx *gin.Context) {

}

func UserDetail(ctx *gin.Context) {
	req := model.SysUser{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	user, err := model.SysUser{}.Detail(req.ID)
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, user)
}

func UserDelete(ctx *gin.Context) {
	req := model.SysUser{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	if err := db.DB().Delete(&req).Error; err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}
