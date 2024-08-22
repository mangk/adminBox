package front

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
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
var frontAdminxJsHandler func(ctx *gin.Context)

func RewriteIndex(f func(ctx *gin.Context)) {
	l := &log.Log{CallerSkip: 0}
	l.SugaredLogger().Infof("RewriteIndex")
	frontIndexHanler = f
}

func RewriteAdminxJs(f func(ctx *gin.Context)) {
	if frontAdminxJsHandler != nil {
		log.Error("frontIndexHanler has been set")
		return
	}
	frontAdminxJsHandler = f
}

func (front) InitModule() {
	root := myHttp.HttpEngine()

	prefix := strings.TrimRight(config.ServerCfg().FrontRouterPrefix, "/")
	assets, _ := fs.Sub(frontFiles, "dist/assets")
	root.StaticFS(prefix+"/assets", http.FS(assets))

	images, _ := fs.Sub(frontFiles, "dist/images")
	root.StaticFS(prefix+"/images", http.FS(images))

	root.GET(prefix+"/", func(ctx *gin.Context) {
		if frontIndexHanler == nil {
			d, _ := fs.ReadFile(frontFiles, "dist/index.html")
			ctx.Data(200, "text/html; charset=utf-8", d)
		} else {
			frontIndexHanler(ctx)
		}
	})

	// 重写 adminx.js
	root.GET(prefix+"/adminx.js", func(ctx *gin.Context) {
		if frontAdminxJsHandler == nil {
			cfg := config.ServerCfg()

			ctx.Data(200, "text/javascript; charset=utf-8", []byte(fmt.Sprintf(`
        window.adminX = {
            Name: '%s',
            RunAt: '%s',
            BackendPrefix: '%s',
            FrontPrefix: '%s',
			Logo: '%s',
			Desc: '%s',
        }
		`, cfg.Name, cfg.RunAt, cfg.BackendRouterPrefix, cfg.FrontRouterPrefix, cfg.Logo, cfg.Desc)))
		} else {
			frontAdminxJsHandler(ctx)
		}
	})

	//admin.SetTmpStr(Convert)
}
