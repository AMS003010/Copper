package main

import (
	"github.com/ams003010/Copper/api-server/initializers"
	"github.com/ams003010/Copper/api-server/routes"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	routes.ImageRegistryRoutes(r)
	r.Run()
}
