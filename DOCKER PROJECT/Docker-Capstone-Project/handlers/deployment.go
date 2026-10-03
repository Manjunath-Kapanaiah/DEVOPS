package handlers

import (
	"net/http"

	"github.com/LikithKumar0112/container-registry-tracker/config"
	"github.com/LikithKumar0112/container-registry-tracker/models"

	"github.com/gin-gonic/gin"
)

func CreateDeployment(c *gin.Context) {

	var deployment models.Deployment

	if err := c.ShouldBindJSON(&deployment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Create(&deployment)

	c.JSON(http.StatusCreated, deployment)
}

func GetDeployments(c *gin.Context) {

	var deployments []models.Deployment

	config.DB.Find(&deployments)

	c.JSON(http.StatusOK, deployments)
}

func GetDeploymentsByEnvironment(c *gin.Context) {

	environment := c.Param("environment")

	var deployments []models.Deployment

	config.DB.Where("environment = ?", environment).Find(&deployments)

	c.JSON(http.StatusOK, deployments)
}
