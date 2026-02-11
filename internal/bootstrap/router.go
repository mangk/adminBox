package bootstrap

import "github.com/gin-gonic/gin"

var globalRouter *gin.Engine

func InitRouter() *gin.Engine {
	globalRouter = gin.Default()
	// Here you can load global middlewares
	return globalRouter
}

func GetRouter() *gin.Engine {
	return globalRouter
}
