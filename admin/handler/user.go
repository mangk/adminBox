package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/admin/model"
	"github.com/mangk/adminBox/db"
	"github.com/mangk/adminBox/http/request"
	"github.com/mangk/adminBox/http/response"
)

func User(ctx *gin.Context) {
	req := request.PublicRequest(ctx)

	var count int64
	list := []model.SysUser{}

	query := db.DB().Model(list)

	if v, has := req.Query["keyword"]; has {
		tv := v.(string)
		if tv != "" {
			tv = "%" + tv + "%"
			query = query.Where("(nick_name LIKE ? OR username LIKE ? OR phone LIKE ? OR email LIKE ?)", tv, tv, tv, tv)
		}
	}
	if v, has := req.Query["enable"]; has {
		tv := v.(string)
		if tv != "" {
			query = query.Where("enable = ?", tv)
		}
	}
	if v, has := req.Query["role"]; has {
		tv := v.(string)
		if tv != "" {
			query = query.Joins("LEFT JOIN "+model.SysUserRole{}.TableName()+" AS sur ON sur.sys_user_id = id").Where("sur.sys_role_id = ?", tv)
		}
	}
	if v, has := req.Query["department"]; has {
		tv := v.(string)
		if tv != "" {
			query = query.Joins("LEFT JOIN "+model.SysUserDepartment{}.TableName()+" AS sud ON sud.sys_user_id = id").Where("sud.sys_department_id = ?", tv)
		}
	}

	if err := query.Count(&count).Error; err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if count > 0 {
		if err := query.Preload("DepartmentList").Preload("RoleList").Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Find(&list).Error; err != nil {
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

	if err := req.Create(); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.OkWithData(ctx, req.ID)
}

func UserEdit(ctx *gin.Context) {
	req := struct {
		model.SysUser
		Password string `json:"password"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	if req.SysUser.Password != "" {
		var salt string
		db.DB().Model(model.SysUser{}).Select("salt").Where("id = ?", req.ID).First(&salt)
		req.Salt = salt
	}

	if err := (model.SysUser{}).Update(req.SysUser); err != nil {
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
	req := struct {
		Old, New, ConfirmNew string
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	u, _ := model.SysUser{}.Detail(request.JWTLoginUserId(ctx))

	_, err := model.SysUser{}.Login(u.Username, req.Old)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}
	if req.New != req.ConfirmNew {
		response.FailWithMsg(ctx, "两次输入的新密码不一致")
		return
	}
	if len(req.New) < 6 {
		response.FailWithMsg(ctx, "新密码不能少于6位")
		return
	}

	u.Password = req.New

	if err := u.Update(u); err != nil {
		response.FailWithError(ctx, err)
		return
	}
	response.Ok(ctx)
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
