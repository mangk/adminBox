package middleware

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

var logger = gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {

	m := map[string]interface{}{
		"status":  p.StatusCode,
		"latency": p.Latency.String(),
		"ip":      p.ClientIP,
		"method":  p.Method,
		"path":    p.Path,
		"errors":  p.ErrorMessage,
	}

	b, _ := json.Marshal(m)
	return string(b) + "\n"
})

func JSONLogger(skipPrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, skipPrefix) {
			c.Next()
			return
		}

		logger(c)
	}
}
