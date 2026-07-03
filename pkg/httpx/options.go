package httpx

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

// ========== 基础选项 ==========

func WithHTTPCode(code int) RespOption {
	return func(r *Resp) {
		if err := validateHTTPCode(code); err != nil {
			r.httpCode = http.StatusInternalServerError
			WithError(err)
			return
		}
		r.httpCode = code
	}
}

func WithCode(code int) RespOption {
	return func(r *Resp) {
		r.Code = code
		if msg, ok := ResponseStatusMap[code]; ok {
			r.Msg = msg
		} else {
			panic(fmt.Errorf("unknown status code: %d", code))
		}
	}
}

func WithMsg(msg string) RespOption {
	return func(r *Resp) {
		msg = strings.Map(func(r rune) rune {
			if r < 32 && r != '\n' && r != '\t' {
				return -1
			}
			return r
		}, msg)
		r.Msg = msg
	}
}

func WithData(data interface{}) RespOption {
	return func(r *Resp) {
		if err := validateDataSize(data); err != nil {
			WithError(err)
			return
		}
		r.Data = data
	}
}

func WithPageInfo(p PageInfo) RespOption {
	return func(r *Resp) {
		r.PageInfo = &p
	}
}

func WithError(err error) RespOption {
	return func(r *Resp) {
		r.Code = -1
		r.error = err
	}
}

// ========== 格式选项 ==========

func AsJSON() RespOption {
	return func(r *Resp) {}
}

func AsString(data string) RespOption {
	return func(r *Resp) {
		r.rawData = data
		r.responseFn = func(ctx *gin.Context) {
			ctx.String(r.httpCode, "%s", data)
		}
	}
}

func AsHTML(templateName string, data interface{}) RespOption {
	return func(r *Resp) {
		if err := validateHTMLTemplate(templateName, data); err != nil {
			WithError(err)
			return
		}

		r.htmlName = templateName
		r.htmlData = data
		r.responseFn = func(ctx *gin.Context) {
			ctx.HTML(r.httpCode, templateName, data)
		}
	}
}

func AsXML(data interface{}) RespOption {
	return func(r *Resp) {
		if err := validateDataSize(data); err != nil {
			WithError(err)
			return
		}

		visited := make(map[uintptr]bool)
		if err := checkHTMLInjectionRecursive(reflect.ValueOf(data), visited); err != nil {
			WithError(err)
			return
		}

		r.Data = data
		r.responseFn = func(ctx *gin.Context) {
			ctx.XML(r.httpCode, data)
		}
	}
}

func AsYAML(data interface{}) RespOption {
	return func(r *Resp) {
		if err := validateDataSize(data); err != nil {
			WithError(err)
			return
		}

		r.Data = data
		r.responseFn = func(ctx *gin.Context) {
			ctx.YAML(r.httpCode, data)
		}
	}
}

func AsSecureJSON(data interface{}) RespOption {
	return func(r *Resp) {
		if err := validateDataSize(data); err != nil {
			WithError(err)
			return
		}

		r.Data = data
		r.responseFn = func(ctx *gin.Context) {
			ctx.SecureJSON(r.httpCode, data)
		}
	}
}

func AsJSONP(data interface{}) RespOption {
	return func(r *Resp) {
		if err := validateDataSize(data); err != nil {
			WithError(err)
			return
		}

		r.Data = data
		r.responseFn = func(ctx *gin.Context) {
			ctx.JSONP(r.httpCode, data)
		}
	}
}

func AsPureJSON(data interface{}) RespOption {
	return func(r *Resp) {
		if err := validateDataSize(data); err != nil {
			WithError(err)
			return
		}

		r.Data = data
		r.responseFn = func(ctx *gin.Context) {
			ctx.PureJSON(r.httpCode, data)
		}
	}
}

func AsAsciiJSON(data interface{}) RespOption {
	return func(r *Resp) {
		if err := validateDataSize(data); err != nil {
			WithError(err)
			return
		}

		r.Data = data
		r.responseFn = func(ctx *gin.Context) {
			ctx.AsciiJSON(r.httpCode, data)
		}
	}
}

// ========== 文件选项 ==========

func AsFile(filePath string) RespOption {
	return func(r *Resp) {
		if err := validateFilePath(filePath); err != nil {
			WithError(err)
			return
		}

		r.filePath = filePath
		r.responseFn = func(ctx *gin.Context) {
			ctx.File(filePath)
		}
	}
}

func AsFileAttachment(filePath, fileName string) RespOption {
	return func(r *Resp) {
		if err := validateFilePath(filePath); err != nil {
			WithError(err)
			return
		}

		if err := validateFilePath(fileName); err != nil {
			WithError(err)
			return
		}

		r.filePath = filePath
		r.fileName = fileName
		r.responseFn = func(ctx *gin.Context) {
			ctx.FileAttachment(filePath, fileName)
		}
	}
}

func AsFileFromFS(filePath string, fs http.FileSystem) RespOption {
	return func(r *Resp) {
		if err := validateFilePath(filePath); err != nil {
			WithError(err)
			return
		}

		if fs == nil {
			WithError(errors.New("file system cannot be nil"))
			return
		}

		r.filePath = filePath
		r.responseFn = func(ctx *gin.Context) {
			ctx.FileFromFS(filePath, fs)
		}
	}
}

// ========== Data 选项 ==========

func AsData(contentType string, data []byte) RespOption {
	return func(r *Resp) {
		if err := validateContentType(contentType); err != nil {
			contentType = "application/octet-stream"
			WithError(err)
			return
		}

		if len(data) > MaxDataSize {
			WithError(ErrDataTooLarge)
			return
		}

		r.contentType = contentType
		r.rawData = data
		r.responseFn = func(ctx *gin.Context) {
			ctx.Data(r.httpCode, contentType, data)
		}
	}
}

func AsDataFromReader(contentType string, reader io.Reader) RespOption {
	var consumed bool
	var mu sync.Mutex

	return func(r *Resp) {
		if err := validateContentType(contentType); err != nil {
			contentType = "application/octet-stream"
			WithError(err)
			return
		}

		if err := validateReader(reader); err != nil {
			WithError(err)
			return
		}

		r.contentType = contentType
		r.reader = reader

		r.responseFn = func(ctx *gin.Context) {
			mu.Lock()
			defer mu.Unlock()

			if consumed {
				ctx.Status(http.StatusInternalServerError)
				ctx.String(http.StatusInternalServerError, "Stream already consumed")
				return
			}
			consumed = true

			ctx.DataFromReader(r.httpCode, -1, contentType, reader, nil)
		}
	}
}

func AsDataFromReaderWithCache(contentType string, reader io.Reader) RespOption {
	return func(r *Resp) {
		if err := validateContentType(contentType); err != nil {
			contentType = "application/octet-stream"
			WithError(err)
			return
		}

		if err := validateReader(reader); err != nil {
			WithError(err)
			return
		}

		var buf bytes.Buffer
		limitedReader := io.LimitReader(reader, MaxDataSize)
		_, err := io.Copy(&buf, limitedReader)
		if err != nil {
			WithError(err)
			return
		}

		cachedData := buf.Bytes()
		if len(cachedData) >= MaxDataSize {
			WithError(ErrDataTooLarge)
			return
		}

		r.cachedData = cachedData
		r.contentLength = len(cachedData)
		r.responseFn = func(ctx *gin.Context) {
			ctx.Data(r.httpCode, contentType, cachedData)
		}
	}
}

func AsDataFromReaderReusable(contentType string, reader io.ReadSeeker) RespOption {
	var mu sync.Mutex

	return func(r *Resp) {
		if err := validateContentType(contentType); err != nil {
			contentType = "application/octet-stream"
			WithError(ErrDataTooLarge)
			return
		}

		if reader == nil {
			WithError(ErrDataTooLarge)
			return
		}

		r.contentType = contentType
		r.reader = reader

		r.responseFn = func(ctx *gin.Context) {
			mu.Lock()
			defer mu.Unlock()

			if seeker, ok := reader.(io.Seeker); ok {
				seeker.Seek(0, io.SeekStart)
			}
			ctx.DataFromReader(r.httpCode, -1, contentType, reader, nil)
		}
	}
}

func WithExtraHeader(key, value string) RespOption {
	return func(r *Resp) {
		if r.extraHeaders == nil {
			r.extraHeaders = make(map[string]string)
		}
		r.extraHeaders[key] = value
	}
}

func WithContentLength(length int) RespOption {
	return func(r *Resp) {
		r.contentLength = length
	}
}
