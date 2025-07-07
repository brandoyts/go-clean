package controller

import (
	"github.com/brandoyts/go-clean/internal/domain"
	"github.com/brandoyts/go-clean/internal/service"
	"github.com/brandoyts/go-clean/internal/utils/httpResponse"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) RegisterUser(c *gin.Context) {
	var requestBody domain.User

	err := c.BindJSON(&requestBody)
	if err != nil {
		httpResponse.Error(c, err)
		return
	}

	err = ac.authService.Register(c, requestBody)
	if err != nil {
		httpResponse.Error(c, err)
		return
	}
}
