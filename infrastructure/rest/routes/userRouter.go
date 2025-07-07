package routes

import (
	"github.com/brandoyts/go-clean/internal/controller"
	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup, handler *controller.UserController) {
	userRouter := r.Group("/users")
	userRouter.GET("", handler.GetAllUsers)
	userRouter.GET("/:id", handler.GetUserById)
	userRouter.POST("/create", handler.CreateUser)
	userRouter.DELETE("/delete/:id", handler.DeleteUser)
}
