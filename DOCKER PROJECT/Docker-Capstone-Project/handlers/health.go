package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/LikithKumar0112/container-registry-tracker/config"

	"github.com/gin-gonic/gin"
)

// HealthCheck reports whether the service is up and can reach the database.
// Returns 200 when healthy and 503 when the database is unreachable, so that
// CI/CD and orchestrators can verify a deployment instead of assuming success.
func HealthCheck(c *gin.Context) {

	sqlDB, err := config.DB.DB()

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "unavailable",
			"database": "error",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":   "unavailable",
			"database": "down",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"database": "up",
	})
}
