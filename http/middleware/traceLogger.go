package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/log"
)

const maxResponseSize = 1 * 1024 * 512 // 0.5 MB

type record struct {
	Ip           string        `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip"`                                   // 请求ip
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:请求方法"`                       // 请求方法
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                             // 请求路径
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:请求状态"`                       // 请求状态
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                            // 代理
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:错误信息"`  // 错误信息
	Header       http.Header   `json:"header" form:"header" gorm:"column:header;serializer:json;comment:请求头"`        // 请求header
	GetParams    url.Values    `json:"get_params" form:"get_params" gorm:"column:get_params;serializer:json;comment:GET参数"`
	Body         string        `json:"body" form:"body" gorm:"type:text;column:body;comment:请求Body"` // 请求Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应Body"` // 响应Body
}

func TraceLogger(flags ...int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		time.Since(start)
		traceLog := log.Trace()
		debug := ctx.Query("debug")
		flag := 0
		if len(flags) > 0 {
			flag = flags[0]
		}

		re := record{
			Ip:     ctx.ClientIP(),
			Method: ctx.Request.Method,
			Path:   ctx.Request.URL.Path,
			Agent:  ctx.Request.UserAgent(),
			Header: ctx.Request.Header,
		}
		// 请求处理
		if !isFileUpload(ctx) &&
			(flag == OperationRecord_LogFlag_Request ||
				flag == OperationRecord_LogFlag_All ||
				debug == fmt.Sprint(OperationRecord_LogFlag_Request) ||
				debug == fmt.Sprint(OperationRecord_LogFlag_All)) {
			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				traceLog.Errorf("read body from request error: %e", err)
			} else {
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}

			re.Body = string(bodyBytes)
			re.GetParams = ctx.Request.URL.Query()
		}

		writer := responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		ctx.Writer = writer
		ctx.Next()

		go func() {
			re.Latency = time.Since(start)
			re.Status = ctx.Writer.Status()
			re.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()

			// 响应处理
			if !isFileDownload(ctx) &&
				(flag == OperationRecord_LogFlag_Response ||
					flag == OperationRecord_LogFlag_All ||
					debug == fmt.Sprint(OperationRecord_LogFlag_Response) ||
					debug == fmt.Sprint(OperationRecord_LogFlag_All)) {
				responseBody := writer.body.String()
				if len(responseBody) > maxResponseSize {
					responseBody = responseBody[:maxResponseSize] // 截断超过 1MB 的部分
				}
				re.Resp = responseBody
			}

			msg, _ := json.Marshal(re)
			traceLog.Info(string(msg))
		}()
	}
}

func isFileUpload(c *gin.Context) bool {
	// 获取Content-Type
	contentType := c.GetHeader("Content-Type")

	// 判断是否是multipart/form-data（常见的文件上传）
	if strings.HasPrefix(contentType, "multipart/form-data") {
		return true
	}

	// 如果是二进制数据上传，考虑Content-Type为application/octet-stream
	if strings.HasPrefix(contentType, "application/octet-stream") {
		return true
	}

	// 如果是其他类型的上传（如二进制文件），可以扩展更多规则
	return false
}

func isFileDownload(c *gin.Context) bool {
	// 检查Content-Disposition头部
	contentDisposition := c.GetHeader("Content-Disposition")

	// 如果包含attachment，通常是文件下载
	if strings.Contains(contentDisposition, "attachment") {
		return true
	}

	// 检查Content-Type，二进制流文件下载
	contentType := c.GetHeader("Content-Type")
	if strings.HasPrefix(contentType, "application/octet-stream") {
		return true
	}

	// 检查是否是常见的文件类型
	fileTypes := []string{
		"application/pdf",
		"image/",
		"audio/",
		"video/",
		"application/vnd.ms-excel", // XLS
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", // XLSX
		"application/vnd.ms-excel.sheet.macroenabled.12",                    // XLSM
	}
	for _, fileType := range fileTypes {
		if strings.HasPrefix(contentType, fileType) {
			return true
		}
	}

	return false
}
