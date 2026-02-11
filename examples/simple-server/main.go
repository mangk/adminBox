package main

import (
	"fmt"
	"github.com/mangk/adminBox/internal/bootstrap"
	"github.com/mangk/adminBox/pkg/config"
	"github.com/mangk/adminBox/pkg/log"

	// Import business modules to register their routes
	// "github.com/mangk/adminBox/internal/module/auth"
	// "github.com/mangk/adminBox/internal/module/user"

	// Anonymous import to trigger the init() function in front/register.go
	_ "github.com/mangk/adminBox/front"
)

// Example server config structure
type ServerConfig struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func main() {
	// 1. Load configuration using the config package
	var cfg ServerConfig
	if err := config.Load("config.yaml", &cfg); err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// 2. Initialize logger using the log package
	log.Init()
	log.Info("Server starting...")

	// 3. Initialize the global router
	router := bootstrap.InitRouter()

	// 4. Register routes from business modules
	// For example:
	// auth.RegisterRoutes(router)
	// user.RegisterRoutes(router)
	// ...

	// 5. Start the server
	port := cfg.Server.Port
	if port == "" {
		port = ":8080" // Default port
	}
	log.Infof("Server listening on port %s", port)
	if err := router.Run(port); err != nil {
		log.Errorf("Server failed to start: %v", err)
	}
}
