package controller

import (
	"github.com/brandoyts/go-clean/internal/domain"
	"github.com/brandoyts/go-clean/internal/service"
	"github.com/brandoyts/go-clean/internal/utils/httpResponse"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{
		userService: svc,
	}
}

func (u *UserController) GetAllUsers(c *gin.Context) {
	result, err := u.userService.GetAllUser(c)
	if err != nil {
		httpResponse.Error(c, err)
		return
	}

	httpResponse.Success(c, result)
}

func (u *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	result, err := u.userService.GetUserById(c, id)
	if err != nil {
		httpResponse.Error(c, err)
		return
	}

	httpResponse.Success(c, result)
}

func (u *UserController) CreateUser(c *gin.Context) {
	var requestBody domain.User
	err := c.BindJSON(&requestBody)
	if err != nil {
		httpResponse.Error(c, err)
	}

	insertedId, err := u.userService.CreateUser(c, requestBody)
	if err != nil {
		httpResponse.Error(c, err)
		return
	}

	httpResponse.Success(c, insertedId)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := u.userService.DeleteUser(c, id)
	if err != nil {
		httpResponse.Error(c, err)
		return
	}

	httpResponse.Success(c, nil)
}
