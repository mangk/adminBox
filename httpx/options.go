package httpx

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ========== 基础选项 ==========

func WithHttpCode(code int) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.httpCode = code
	}
}

func WithCode(code int) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.Code = code
		if msg, ok := ResponseStatusMap[code]; ok {
			r.Msg = msg
		}
	}
}

func WithMsg(msg string) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.Msg = msg
	}
}

func WithData(data any) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.Data = data
	}
}

func WithPageInfo(p PageInfo) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.PageInfo = &p
	}
}

func WithError(err error, code ...int) RespOption {
	return func(c *gin.Context, r *Resp) {
		if len(code) == 1 {
			r.Code = code[0]
		} else {
			r.Code = -1
		}
		r.error = err
	}
}

func WithCustomerSet(fn func(c *gin.Context, r *Resp)) RespOption {
	return fn
}

// ========== 格式选项 ==========
func AsInnerJson() RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.JSON(r.httpCode, *r)
		})
	}
}

func AsJSON() RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.JSON(r.httpCode, r.Data)
		})
	}
}

func AsString() RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.String(r.httpCode, "%s", r.Data)
		})
	}
}

func AsHTML(templateName string) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.HTML(r.httpCode, templateName, r.Data)
		})
	}
}

func AsXML() RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.XML(r.httpCode, r.Data)
		})
	}
}

func AsYAML() RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.YAML(r.httpCode, r.Data)
		})
	}
}

func AsSecureJSON() RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.SecureJSON(r.httpCode, r.Data)
		})
	}
}

func AsJSONP() RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.JSONP(r.httpCode, r.Data)
		})
	}
}

func AsPureJSON() RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.PureJSON(r.httpCode, r.Data)
		})
	}
}

func AsAsciiJSON() RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.AsciiJSON(r.httpCode, r.Data)
		})
	}
}

func AsFile(filePath string) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.File(filePath)
		})
	}
}

func AsFileAttachment(filePath, fileName string) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.FileAttachment(filePath, fileName)
		})
	}
}

func AsFileFromFS(filePath string, fs http.FileSystem) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.FileFromFS(filePath, fs)
		})
	}
}

func AsData(contentType string) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.Data(r.httpCode, contentType, r.Data.([]byte))
		})
	}
}

func AsDataFromReader(contentType string, reader io.Reader) RespOption {
	return func(c *gin.Context, r *Resp) {
		r.SetResponseFunc(func(ctx *gin.Context) {
			ctx.DataFromReader(r.httpCode, -1, contentType, reader, nil)
		})
	}
}

func AsCustomerSet(fn func(c *gin.Context, r *Resp)) RespOption {
	return fn
}
