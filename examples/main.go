package main

import (
	_ "github.com/mangk/adminBox/pkg/admin/api"
	_ "github.com/mangk/adminBox/pkg/admin/front"
	"github.com/mangk/adminBox/pkg/httpServer"
)

func main() {
	httpServer.Execute("simple_server", "examples")
}
