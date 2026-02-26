package main

import (
	_ "github.com/mangk/adminBox/pkg/admin/api"
	_ "github.com/mangk/adminBox/pkg/admin/front"
	"github.com/mangk/adminBox/pkg/httpServer"
)

// Example server config structure
type ServerConfig struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func main() {
	httpServer.Execute("simple_server", "examples")
}
