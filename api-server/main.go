package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ams003010/Copper/api-server/initializers"
	"github.com/ams003010/Copper/api-server/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	routes.ImageRegistryRoutes(r)
	r.Run()
}
