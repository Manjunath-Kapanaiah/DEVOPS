package handlers

import (
	"net/http"

	"github.com/LikithKumar0112/container-registry-tracker/config"
	"github.com/LikithKumar0112/container-registry-tracker/models"

	"github.com/gin-gonic/gin"
)

func CreateRegistry(c *gin.Context) {

	var registry models.Registry

	if err := c.ShouldBindJSON(&registry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Create(&registry)

	c.JSON(http.StatusCreated, registry)
}

func GetRegistries(c *gin.Context) {

	var registries []models.Registry

	config.DB.Find(&registries)

	c.JSON(http.StatusOK, registries)
}
