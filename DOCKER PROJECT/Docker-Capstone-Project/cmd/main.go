package main

import (
	"github.com/LikithKumar0112/container-registry-tracker/config"
	"github.com/LikithKumar0112/container-registry-tracker/models"
	"github.com/LikithKumar0112/container-registry-tracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.Registry{},
		&models.Image{},
		&models.Deployment{},
	)

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8081")
}
