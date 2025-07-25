package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheckRouter(r *gin.RouterGroup) {
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})
}
