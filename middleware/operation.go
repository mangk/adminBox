package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/mangk/adminBox/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

const OperationRecord_LogFlag_Request = 1
const OperationRecord_LogFlag_Response = 2
const OperationRecord_LogFlag_All = 3

// 记录日志的中间件
// flag 1:记录请求日志;2:记录响应日志;3:记录请求和响应;
func OperationRecord(flag int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				log.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}

		record := struct {
			Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`                                   // 请求ip
			Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法"`                       // 请求方法
			Path         string        `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                             // 请求路径
			Status       int           `json:"status" form:"status" gorm:"column:status;comment:请求状态"`                       // 请求状态
			Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
			Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                            // 代理
			ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`  // 错误信息
			Header       http.Header   `json:"header" form:"header" gorm:"column:header;serializer:json;"`                   // 请求header
			Body         string        `json:"body" form:"body" gorm:"type:text;column:body;comment:请求Body"`                 // 请求Body
			Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应Body"`                 // 响应Body
		}{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Header: c.Request.Header,
			Agent:  c.Request.UserAgent(),
		}
		if flag == OperationRecord_LogFlag_Request || flag == OperationRecord_LogFlag_All {
			record.Body = string(body)
		}

		// 上传文件时候 中间件日志进行裁断操作
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			if len(record.Body) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Body)
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		if flag == OperationRecord_LogFlag_Response || flag == OperationRecord_LogFlag_All {
			record.Resp = writer.body.String()
		}

		if strings.Contains(c.Writer.Header().Get("Pragma"), "public") ||
			strings.Contains(c.Writer.Header().Get("Expires"), "0") ||
			strings.Contains(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/force-download") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/download") ||
			strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment") ||
			strings.Contains(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
			if len(record.Resp) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Resp)
				record.Resp = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		msg, _ := json.Marshal(record)
		log.Info(string(msg))
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
