package routes

import (
	"github.com/ams003010/Copper/api-server/handlers"
	"github.com/gin-gonic/gin"
)

func ImageRegistryRoutes(r *gin.Engine) {
	imageGroup := r.Group("/images")
	{
		// Image Routes
		imageGroup.POST("", handlers.ImageCreate)            // Create an image
		imageGroup.GET("", handlers.GetAllImages)            // Read all images
		imageGroup.PUT("/:image/:tag", handlers.UpdateImage) // Update an image
		imageGroup.DELETE("/:image", handlers.DeleteImage)   // Delete an image
	}
}
