package front

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mangk/adminBox/pkg/httpServer"
)

//go:embed all:dist
var staticFiles embed.FS

//go:embed crudTemplate.vue
var CrudTemplate string

func init() {
	httpServer.SetRouter(func(root *gin.Engine) {
		dist, err := fs.Sub(staticFiles, "dist")
		if err != nil {
			panic(err)
		}
		root.StaticFS("/", http.FS(dist))
	})
}
