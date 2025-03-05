package main

import (
	"github.com/ams003010/Copper/api-server/initializers"
	"github.com/ams003010/Copper/api-server/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.RegistryImage{})
}
