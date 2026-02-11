package front

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/mangk/adminBox/internal/bootstrap"
)

//go:embed all:dist
var staticFiles embed.FS

func init() {
	router := bootstrap.GetRouter()
	if router == nil {
		// This can happen if the router is not initialized before this package is imported.
		// It's a good practice to handle this case.
		return
	}
	
	dist, err := fs.Sub(staticFiles, "dist")
	if err != nil {
		panic(err)
	}

	// Serve the static files under /
	// The root path will serve index.html
	router.StaticFS("/", http.FS(dist))

	// Optional: If you want to redirect a specific path like /ui to the root
	// router.GET("/ui/*filepath", func(c *gin.Context) {
	// 	c.FileFromFS(c.Param("filepath"), http.FS(dist))
	// })
	// router.GET("/ui", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "/ui/index.html")
	// })
}
