package routes

import (
	"github.com/LikithKumar0112/container-registry-tracker/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/health", handlers.HealthCheck)

	router.POST("/registries", handlers.CreateRegistry)
	router.GET("/registries", handlers.GetRegistries)
	
	router.POST("/images", handlers.CreateImage)
	router.GET("/images", handlers.GetImages)
	router.GET("/images/:id", handlers.GetImageByID)

	router.POST("/deployments", handlers.CreateDeployment)
	router.GET("/deployments", handlers.GetDeployments)
	router.GET("/deployments/environment/:environment",
	handlers.GetDeploymentsByEnvironment)
}
