package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/admin/model"
	"github.com/mangk/adminX/cache"
	"github.com/mangk/adminX/config"
	"github.com/mangk/adminX/db"
	"github.com/mangk/adminX/front"
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
		response.FailWithError(ctx, err)
		return
	}

	if !(cache.Base64CaptchaStore{}.Verify(req.CaptchaId, req.VerificationCode, true)) {
		response.FailWithMsg(ctx, "验证码错误")
		return
	}

	user, err := model.SysUser{}.Login(req.Username, req.Password)
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}

	jwtToken, err := model.NewJWT([]byte(config.JwtCfg().SigningKey)).Create(user, config.JwtCfg())
	if err != nil {
		response.FailWithError(ctx, err)
		return
	}

	user.JwtToken = jwtToken

	if err := (model.SysCasbinRole{}).UpdateCasbin(user.ID); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	response.OkWithData(ctx, user)
}

func AuthLogout(ctx *gin.Context) {

}

func IsRewriteIndex(ctx *gin.Context) {
	response.OkWithData(ctx, front.IsRewriteIndex())
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
		response.FailWithError(ctx, err)
		return
	}

	response.OkWithData(ctx, map[string]interface{}{
		"menu": tree,
		"user": user,
	})
}

func AuthPermissionAll(ctx *gin.Context) {
	list := []Item{}

	for _, menu := range (model.SysMenu{}).All(true) {
		list = append(list, Item{
			ID:     fmt.Sprintf("menu%s%d", split, menu.ID),
			Pid:    fmt.Sprintf("menu%s%d", split, menu.Pid),
			Name:   menu.Meta.Title,
			Type:   "menu",
			SelfID: menu.ID,
		})
	}

	for _, api := range (model.SysApi{}).All(true) {
		list = append(list, Item{
			ID:     fmt.Sprintf("api%s%d", split, api.ID),
			Pid:    fmt.Sprintf("menu%s%d", split, api.MenuId),
			Name:   api.Name,
			Type:   "api",
			SelfID: api.ID,
		})
	}

	d, _ := Item{}.buildTree(list)

	root := Item{
		Name:     "根目录",
		Children: d,
	}
	withRoot := []Item{}
	withRoot = append(withRoot, root)
	response.OkWithData(ctx, withRoot)
}

func AuthPermissionGetByIdAndModule(ctx *gin.Context) {
	req := struct {
		Id     int
		Module string
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithError(ctx, err)
		return
	}

	curSet, resultSet, otherSetList, err := (model.SysAuth{}).LoadDetail(req.Id, req.Module)
	if err != nil {
		response.FailWithError(ctx, err)
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
		response.FailWithError(ctx, err)
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
		response.FailWithError(ctx, err)
		return
	}

	response.Ok(ctx)
}

var split = "|"

type Item struct {
	ID       string `json:"id"`
	SelfID   int    `json:"selfId"`
	Pid      string `json:"pid"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Children []Item `json:"children"`
}

func (s Item) buildTree(menuList []Item) ([]Item, error) {
	var err error
	treeMap := make(map[string][]Item)
	for _, v := range menuList {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	resMenuList := treeMap["menu"+split+"0"]
	for i := 0; i < len(resMenuList); i++ {
		err = s.getBaseChildrenList(&resMenuList[i], treeMap)
	}
	return resMenuList, err
}

func (s Item) getBaseChildrenList(menu *Item, treeMap map[string][]Item) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = s.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
