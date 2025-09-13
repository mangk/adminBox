package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/admin/model"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/http/request"
	"github.com/mangk/adminBox/http/response"
)

func JWTCheckByCasbin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" || len(token) <= 6 {
			response.FailWithCodeAndNeedReload(ctx, http.StatusUnauthorized, "无效Token")
			ctx.Abort()
			return
		}

		jwtUserInfo, err := model.NewJWT([]byte(config.JwtCfg().SigningKey)).Parse(token[7:])
		if err != nil || jwtUserInfo == nil {
			response.FailWithCodeAndNeedReload(ctx, http.StatusUnauthorized, "Token解析失败")
			ctx.Abort()
			return
		}

		if time.Now().Unix() > jwtUserInfo.ExpiresAt {
			response.FailWithCodeAndNeedReload(ctx, http.StatusUnauthorized, "授权已过期")
			ctx.Abort()
			return
		}

		ctx.Set(request.ContextLoginUserKey, jwtUserInfo.UserId)
		ctx.Set(request.ContextRoleUserTypeKey, ctx.Request.Header.Get("X-User-Type"))
		roleUserId, _ := strconv.Atoi(ctx.Request.Header.Get("X-User-Id"))
		ctx.Set(request.ContextRoleUserIdKey, roleUserId)

		sub := jwtUserInfo.Id
		obj := ctx.Request.URL.Path
		act := ctx.Request.Method

		adminPrefix := config.ServerCfg().BackendRouterPrefix
		if adminPrefix != "" {
			objs := strings.Split(obj, adminPrefix)
			obj = objs[len(objs)-1]
		}

		access, err := model.LoadEnforce().Enforce(sub, obj, act)
		if err != nil {
			response.FailWithError(ctx, err)
			ctx.Abort()
			return
		}
		if !access {
			response.FailWithCodeAndNeedReload(ctx, http.StatusUnauthorized, "未授权访问")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
