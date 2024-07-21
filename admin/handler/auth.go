package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/admin/model"
	"github.com/mangk/adminX/cache"
	"github.com/mangk/adminX/config"
	"github.com/mangk/adminX/db"
	"github.com/mangk/adminX/http/request"
	"github.com/mangk/adminX/http/response"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
)

func AuthLogin(ctx *gin.Context) {
	req := struct {
		CaptchaId        string `json:"captcha_id" binding:"required"`
		Username         string `json:"username" binding:"required"`
		Password         string `json:"password" binding:"required"`
		VerificationCode string `json:"verification_code" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	if !(cache.Base64CaptchaStore{}.Verify(req.CaptchaId, req.VerificationCode, true)) {
		response.FailWithMsg(ctx, "验证码错误")
		return
	}

	user, err := model.SysUser{}.Login(req.Username, req.Password)
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	jwtToken, err := model.NewJWT([]byte(config.JwtCfg().SigningKey)).Create(user, config.JwtCfg())
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	user.JwtToken = jwtToken

	if err := (model.SysCasbinRole{}).UpdateCasbin(user.ID); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, user)
}

func AuthLogout(ctx *gin.Context) {

}

func AuthVerificationCode(ctx *gin.Context) {
	cfg := config.CaptchaCfg()
	driver := base64Captcha.NewDriverDigit(cfg.ImgHeight, cfg.ImgWidth, cfg.KeyLong, 0.7, 80)

	cp := base64Captcha.NewCaptcha(driver, cache.Base64CaptchaStore{})
	id, b64s, _, err := cp.Generate()
	if err != nil {
		response.FailWithMsg(ctx, "验证码获取失败")
		return
	}
	response.OkWithDetail(ctx, "验证码获取成功", map[string]interface{}{
		"captcha_id": id,
		"pic_path":   b64s,
	})
}

func AuthUserPermission(ctx *gin.Context) {
	user := request.JWTLoginUserFetch(ctx)
	tree, err := (model.SysMenu{}).Tree(true, false, true, user.ID)
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, map[string]interface{}{
		"menu": tree,
		"user": user,
	})
}

func AuthPermissionAll(ctx *gin.Context) {
	if data, err := (model.SysMenu{}).Tree(true, true, true); err != nil {
		response.FailWithMsg(ctx, err.Error())
	} else {
		response.OkWithData(ctx, data)
	}
}

func AuthPermissionGetByIdAndModule(ctx *gin.Context) {
	req := struct {
		Id     int
		Module string
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	curSet, resultSet, otherSetList, err := (model.SysAuth{}).LoadDetail(req.Id, req.Module)
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	response.OkWithData(ctx, gin.H{
		"cur_set":        curSet,
		"result_set":     resultSet,
		"ohter_set_list": otherSetList,
	})
}

func AuthPermissionSave(ctx *gin.Context) {
	req := struct {
		Id     int
		Module string
		List   map[string]int
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	allow := []string{model.SysUser{}.TableName(), model.SysRole{}.TableName(), model.SysDepartment{}.TableName()}
	var canSave bool
	for _, st := range allow {
		if st == req.Module {
			canSave = true
		}
	}
	if !canSave {
		response.FailWithMsg(ctx, "模块类型不支持")
		return
	}

	save := []model.SysAuth{}
	for k, v := range req.List {
		typeAndKey := strings.Split(k, model.SysAuth{}.SplitStr())
		if len(typeAndKey) != 2 {
			continue
		}
		if v == 0 {
			continue
		}
		save = append(save, model.SysAuth{
			TableId:     req.Id,
			TableModule: req.Module,
			Type:        typeAndKey[0],
			Key:         typeAndKey[1],
			SetValue:    v,
		})
	}

	err := db.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("table_id = ? AND table_module = ?", req.Id, req.Module).Delete(&model.SysAuth{}).Error; err != nil {
			return err
		}

		if len(save) > 0 {
			if err := tx.Create(&save).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		response.FailWithMsg(ctx, err.Error())
		return
	}

	response.Ok(ctx)
}
