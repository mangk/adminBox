package front

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox"
	"github.com/mangk/adminBox/config"
	"github.com/mangk/adminBox/log"

	_ "github.com/mangk/adminBox/admin"
)

//go:embed dist
var frontFiles embed.FS

//go:embed crudTemplate.vue
var CrudTemplate string

func init() {
	adminBox.SetRouter(func(root *gin.Engine) {
		path := ""
		frontPrefix := config.ServerCfg().FrontRouterPrefix
		if frontPrefix != "" {
			path = strings.TrimRight(frontPrefix, "/")
		}
		root.GET(path, func(ctx *gin.Context) {
			_, has := ctx.GetQuery("front")
			if frontIndexHanler != nil && !has {
				frontIndexHanler(ctx)
			} else {
				d, _ := fs.ReadFile(frontFiles, "dist/index.html")
				ctx.Data(200, "text/html; charset=utf-8", d)
			}
		})

		assets, _ := fs.Sub(frontFiles, "dist/assets")
		root.StaticFS("/assets", http.FS(assets))

		images, _ := fs.Sub(frontFiles, "dist/images")
		root.StaticFS("/images", http.FS(images))

		// 重写 adminBox.js
		root.SetHTMLTemplate(template.Must(template.New("adminBox.js").Delims("/***", "***/").ParseFS(frontFiles, "dist/adminBox.js")))
		root.GET("/adminBox.js", func(ctx *gin.Context) {
			ctx.Header("Content-Type", "application/javascript")
			allConfig := fmt.Sprintf(`BackendRouterPrefix: '%s',%s`, config.ServerCfg().BackendRouterPrefix, writeByadminBoxConfig)
			ctx.HTML(http.StatusOK, "adminBox.js", gin.H{
				"writeByadminBox_config": template.HTML(allConfig),
				"writeByadminBox_func":   template.HTML(writeByadminBoxFunc),
			})
		})
	})

}

var frontIndexHanler func(ctx *gin.Context)

var writeByadminBoxConfig string
var writeByadminBoxFunc string

func RewriteIndex(f func(ctx *gin.Context)) {
	log.Info("RewriteIndex")
	frontIndexHanler = f
}

func SetAdminBoxJsUserCodeSnippet(cfg, function string) {
	writeByadminBoxConfig = cfg
	writeByadminBoxFunc = function
}

func DynamicTemplate(ctx *gin.Context, filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "File reading error: %s", err.Error())
		return
	}
	// 发送原始 HTML 内容
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", content)
}

func TemplateBuild(filePath, embedStr string) func(ctx *gin.Context) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return func(ctx *gin.Context) {
			ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(embedStr))
		}
	} else {
		return func(ctx *gin.Context) {
			content, err := os.ReadFile(filePath)
			if err != nil {
				ctx.String(http.StatusInternalServerError, "File reading error: %s", err.Error())
				return
			}
			// 发送原始 HTML 内容
			ctx.Data(http.StatusOK, "text/html; charset=utf-8", content)
		}
	}
}
