package handlers

import (
	"net/http"

	"github.com/LikithKumar0112/container-registry-tracker/config"
	"github.com/LikithKumar0112/container-registry-tracker/models"

	"github.com/gin-gonic/gin"
)

func CreateImage(c *gin.Context) {

	var image models.Image

	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Create(&image)

	c.JSON(http.StatusCreated, image)
}

func GetImages(c *gin.Context) {

	var images []models.Image

	config.DB.Find(&images)

	c.JSON(http.StatusOK, images)
}

func GetImageByID(c *gin.Context) {

	var image models.Image

	id := c.Param("id")

	if err := config.DB.First(&image, id).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Image not found",
		})

		return
	}

	c.JSON(http.StatusOK, image)
}
