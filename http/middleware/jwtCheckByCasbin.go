package middleware

import (
	"net/http"
	"strconv"
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
			response.FailWithCode(ctx, http.StatusUnauthorized, "鉴权失败")
			ctx.Abort()
			return
		}

		jwtUserInfo, err := model.NewJWT([]byte(config.JwtCfg().SigningKey)).Parse(token[7:])
		if err != nil || jwtUserInfo == nil {
			response.FailWithCode(ctx, http.StatusUnauthorized, "鉴权失败")
			ctx.Abort()
			return
		}

		if time.Now().Unix() > jwtUserInfo.ExpiresAt {
			response.FailWithCode(ctx, http.StatusUnauthorized, "授权已过期")
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

		access, err := model.LoadEnforce().Enforce(sub, obj, act)
		if err != nil {
			response.FailWithError(ctx, err)
			ctx.Abort()
			return
		}
		if !access {
			response.FailWithCode(ctx, http.StatusUnauthorized, "鉴权失败")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
