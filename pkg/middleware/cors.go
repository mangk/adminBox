package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/pkg/config"
	"github.com/mangk/adminBox/pkg/log"
	"go.uber.org/zap"
)

func Cors(why string) gin.HandlerFunc {
	log.Zaplog().WithOptions(zap.AddCallerSkip(-1)).Sugar().Infof("use middleware cors: %s", why)
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, X-User-Type, X-User-Id")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}

func CorsByRules() gin.HandlerFunc {
	// 放行全部
	if config.CORSCfg().Mode == "allowAll" {
		return Cors("allowAll")
	}
	return func(c *gin.Context) {
		whitelist := checkCors(c.GetHeader("origin"))

		// 通过检查, 添加请求头
		if whitelist != nil {
			c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		// 严格白名单模式且未通过检查，直接拒绝处理请求
		if whitelist == nil && config.CORSCfg().Mode == "strictWhitelist" {
			c.AbortWithStatus(http.StatusForbidden)
		}

		// 处理请求
		c.Next()
	}
}

func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range config.CORSCfg().Whitelist {
		// 遍历配置中的跨域头，寻找匹配项
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
