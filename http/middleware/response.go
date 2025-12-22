package middleware

import (
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/http/response"
	"github.com/mangk/adminBox/log"
)

func ResponseRecover() gin.HandlerFunc {
	if !response.IsUseResponseRecover() {
		return gin.Recovery()
	}
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(response.ResponseAbort); ok {
					return
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				headers := strings.Split(string(httpRequest), "\r\n")
				for idx, header := range headers {
					current := strings.Split(header, ":")
					if current[0] == "Authorization" {
						headers[idx] = current[0] + ": *"
					}
				}
				headersToStr := strings.Join(headers, "\r\n")
				log.Errorf("[Recovery] %s panic recovered:\n%s\n%v\n%s\n",
					time.Now().Format("2006/01/02 - 15:04:05"),
					headersToStr,
					r,
					debug.Stack())

				if !c.Writer.Written() {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}
