package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/admin/model"
	"github.com/mangk/adminX/config"
	"github.com/mangk/adminX/http/request"
	"github.com/mangk/adminX/http/response"
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

		ctx.Set(request.ContextUserKey, jwtUserInfo.UserId)

		sub := jwtUserInfo.Id
		obj := ctx.Request.URL.Path
		act := ctx.Request.Method

		access, err := model.LoadEnforce().Enforce(sub, obj, act)
		if err != nil {
			response.FailWithMsg(ctx, err.Error())
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
