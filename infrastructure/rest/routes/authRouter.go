package routes

import (
	"github.com/brandoyts/go-clean/internal/controller"
	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup, handler *controller.AuthController) {
	authRouter := r.Group("/auth")
	authRouter.POST("/register", handler.RegisterUser)
}
