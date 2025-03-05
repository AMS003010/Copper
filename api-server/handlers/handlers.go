package handlers

import (
	"net/http"

	"github.com/ams003010/Copper/api-server/initializers"
	"github.com/ams003010/Copper/api-server/models"
	"github.com/gin-gonic/gin"
)

func ImageCreate(c *gin.Context) {
	var body struct {
		Image     string `json:"image"`
		Tag       string `json:"tag"`
		Timestamp string `json:"timestamp"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var existingImage models.RegistryImage
	result := initializers.DB.Where("image = ? AND tag = ?", body.Image, body.Tag).First(&existingImage)

	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Image with this tag already exists"})
		return
	}

	image := models.RegistryImage{
		Image:     body.Image,
		Tag:       body.Tag,
		Timestamp: body.Timestamp,
	}

	if err := initializers.DB.Create(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save image"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Image saved successfully", "image": image})
}

func GetAllImages(c *gin.Context) {
	var images []models.RegistryImage
	result := initializers.DB.Find(&images)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching images"})
		return
	}

	if len(images) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No images found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"images": images})
}

func UpdateImage(c *gin.Context) {
	image := c.Param("image")
	tag := c.Param("tag")

	// Check if parameters are empty
	if image == "" || tag == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid image and tag field"})
		return
	}

	// Find the image in the database
	var imageToUpdate models.RegistryImage
	result := initializers.DB.Where("image = ? AND tag = ?", image, tag).First(&imageToUpdate)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	// Parse the request body
	var updateData struct {
		Image     string `json:"image"`
		Tag       string `json:"tag"`
		Timestamp string `json:"timestamp"`
	}

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updates := make(map[string]interface{})
	if updateData.Image != "" {
		updates["image"] = updateData.Image
	}
	if updateData.Tag != "" {
		updates["tag"] = updateData.Tag
	}
	if updateData.Timestamp != "" {
		updates["timestamp"] = updateData.Timestamp
	}

	initializers.DB.Model(&imageToUpdate).Updates(updates)

	c.JSON(http.StatusOK, gin.H{"message": "Image updated successfully", "updated_image": imageToUpdate})
}

func DeleteImage(c *gin.Context) {
	image := c.Param("image")

	if image == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid image field"})
		return
	}

	var foundImage models.RegistryImage
	result := initializers.DB.Where("image = ?", image).First(&foundImage)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No image like that exists"})
		return
	}

	initializers.DB.Delete(&foundImage)

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully", "deleted_image": foundImage})
}
