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
	"github.com/mangk/adminX/moduleRegister"
)

//go:embed dist
var frontFiles embed.FS

func init() {
	moduleRegister.ModuleAdd(front{})
}

type front struct{}

func (front) InitModule() {
	root := myHttp.HttpEngine()

	prefix := strings.TrimRight(config.ServerCfg().FrontRouterPrefix, "/")
	assets, _ := fs.Sub(frontFiles, "dist/assets")
	root.StaticFS(prefix+"/assets", http.FS(assets))

	images, _ := fs.Sub(frontFiles, "dist/images")
	root.StaticFS(prefix+"/images", http.FS(images))

	root.GET(prefix+"/", func(ctx *gin.Context) {
		d, _ := fs.ReadFile(frontFiles, "dist/index.html")
		ctx.Data(200, "text/html; charset=utf-8", d)
	})

	// 重写 adminx.js
	root.GET(prefix+"/adminx.js", func(ctx *gin.Context) {
		cfg := config.ServerCfg()

		ctx.Data(200, "text/javascript; charset=utf-8", []byte(fmt.Sprintf(`
        window.adminX = {
            Name: '%s',
            RunAt: '%s',
            BackendPrefix: '%s',
            FrontPrefix: '%s',
        }
		`, cfg.Name, cfg.RunAt, cfg.BackendRouterPrefix, cfg.FrontRouterPrefix)))
	})

	//admin.SetTmpStr(Convert)
}
