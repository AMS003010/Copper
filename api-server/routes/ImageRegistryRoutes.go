package routes

import (
	"github.com/ams003010/Copper/api-server/handlers"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func ImageRegistryRoutes(r *gin.Engine, redisClient *redis.Client) {
	imageGroup := r.Group("/images")
	{
		// Image Routes
		imageGroup.POST("", handlers.ImageCreate) // Create an image
		imageGroup.GET("", func(c *gin.Context) {
			handlers.GetAllImages(c, redisClient)
		}) // Read all images
		imageGroup.PUT("/:image/:tag", handlers.UpdateImage) // Update an image
		imageGroup.DELETE("/:image", handlers.DeleteImage)   // Delete an image
	}
}
