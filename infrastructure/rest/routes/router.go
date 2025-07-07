package routes

import (
	"github.com/brandoyts/go-clean/internal/controller"
	"github.com/gin-gonic/gin"
)

func Initialize(
	userController *controller.UserController,
	authController *controller.AuthController,
) *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/api/v1")

	// register routes
	healthCheckRouter(apiRouter)
	userRouter(apiRouter, userController)
	authRouter(apiRouter, authController)

	return router
}
