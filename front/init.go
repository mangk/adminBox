package front

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/config"
	myHttp "github.com/mangk/adminBox/http"
	"github.com/mangk/adminBox/log"
	"github.com/mangk/adminBox/moduleRegister"

	_ "github.com/mangk/adminBox/admin"
)

//go:embed dist
var frontFiles embed.FS

func init() {
	moduleRegister.ModuleAdd(front{})
}

type front struct{}

var frontIndexHanler func(ctx *gin.Context)

var writeByadminBoxConfig string
var writeByadminBoxFunc string

func RewriteIndex(f func(ctx *gin.Context)) {
	l := &log.Log{CallerSkip: 0}
	l.SugaredLogger().Infof("RewriteIndex")
	frontIndexHanler = f
}

func IsRewriteIndex() bool {
	return frontIndexHanler != nil
}

func SetadminBoxJsUserCodeSnippet(cfg, function string) {
	writeByadminBoxConfig = cfg
	writeByadminBoxFunc = function
}

func LoadIndexPathPrefix() string {
	indexPathPrefix := ""
	if config.ServerCfg().FrontRouterPrefix != "" {
		indexPathPrefix = strings.TrimRight(config.ServerCfg().FrontRouterPrefix, "/")
	}
	return indexPathPrefix
}

func (front) InitModule() {
	root := myHttp.HttpEngine()

	indexPathPrefix := LoadIndexPathPrefix()

	if indexPathPrefix != "" {
		root.GET(indexPathPrefix+"/", func(ctx *gin.Context) {
			d, _ := fs.ReadFile(frontFiles, "dist/index.html")
			ctx.Data(200, "text/html; charset=utf-8", d)
		})
	}

	if indexPathPrefix == "" {
		root.GET("/", func(ctx *gin.Context) {
			if frontIndexHanler != nil {
				frontIndexHanler(ctx)
			} else {
				d, _ := fs.ReadFile(frontFiles, "dist/index.html")
				ctx.Data(200, "text/html; charset=utf-8", d)
			}
		})

	}

	assets, _ := fs.Sub(frontFiles, "dist/assets")
	root.StaticFS("/assets", http.FS(assets))

	images, _ := fs.Sub(frontFiles, "dist/images")
	root.StaticFS("/images", http.FS(images))

	// 重写 adminBox.js
	root.SetHTMLTemplate(template.Must(template.New("").Delims("/***", "***/").ParseFS(frontFiles, "dist/adminBox.js")))
	root.GET("/adminBox.js", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/javascript")
		ctx.HTML(http.StatusOK, "adminBox.js", gin.H{
			"writeByadminBox_config": template.HTML(writeByadminBoxConfig),
			"writeByadminBox_func":   template.HTML(writeByadminBoxFunc),
		})
	})

	//admin.SetTmpStr(Convert)
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
