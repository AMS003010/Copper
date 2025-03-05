package main

import (
	"github.com/ams003010/Copper/api-server/initializers"
	"github.com/ams003010/Copper/api-server/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	routes.ImageRegistryRoutes(r)
	r.Run()
}
