package front

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminX/config"
	myHttp "github.com/mangk/adminX/http"
	"github.com/mangk/adminX/log"
	"github.com/mangk/adminX/moduleRegister"
)

//go:embed dist
var frontFiles embed.FS

func init() {
	moduleRegister.ModuleAdd(front{})
}

type front struct{}

var frontIndexHanler func(ctx *gin.Context)

var writeByAdminXConfig string
var writeByAdminXFunc string

func RewriteIndex(f func(ctx *gin.Context)) {
	l := &log.Log{CallerSkip: 0}
	l.SugaredLogger().Infof("RewriteIndex")
	frontIndexHanler = f
}

func IsRewriteIndex() bool {
	return frontIndexHanler != nil
}

func SetAdminxJsUserCodeSnippet(cfg, function string) {
	writeByAdminXConfig = cfg
	writeByAdminXFunc = function
}

func LoadIndexPathPrefix() string {
	indexPathPrefix := "_"
	if config.ServerCfg().FrontRouterPrefix != "" {
		indexPathPrefix = strings.TrimRight(config.ServerCfg().FrontRouterPrefix, "/")
	}
	return indexPathPrefix
}

func (front) InitModule() {
	root := myHttp.HttpEngine()

	indexPathPrefix := LoadIndexPathPrefix()

	root.GET("/", func(ctx *gin.Context) {
		// TODO 使用重写后需要前端页面感知，不在展示登录页面，而是404 之类的
		if frontIndexHanler != nil {
			frontIndexHanler(ctx)
		} else {
			d, _ := fs.ReadFile(frontFiles, "dist/index.html")
			ctx.Data(200, "text/html; charset=utf-8", d)
		}
	})

	root.GET(indexPathPrefix+"/", func(ctx *gin.Context) {
		d, _ := fs.ReadFile(frontFiles, "dist/index.html")
		ctx.Data(200, "text/html; charset=utf-8", d)
	})

	assets, _ := fs.Sub(frontFiles, "dist/assets")
	root.StaticFS("/assets", http.FS(assets))

	images, _ := fs.Sub(frontFiles, "dist/images")
	root.StaticFS("/images", http.FS(images))

	// 重写 adminx.js
	root.SetHTMLTemplate(template.Must(template.New("").Delims("/***", "***/").ParseFS(frontFiles, "dist/adminx.js")))
	root.GET("/adminx.js", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/javascript")
		ctx.HTML(http.StatusOK, "adminx.js", gin.H{
			"writeByAdminX_config": template.HTML(writeByAdminXConfig),
			"writeByAdminX_func":   template.HTML(writeByAdminXFunc),
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
