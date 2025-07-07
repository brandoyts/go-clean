package httpResponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	errorMessage := "error"

	if err != nil {
		errorMessage = err.Error()
	}

	c.JSON(http.StatusOK, response{
		Message: errorMessage,
		Code:    http.StatusBadRequest,
		Data:    nil,
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusOK, response{
		Message: "Unauthorized",
		Code:    http.StatusUnauthorized,
		Data:    nil,
	})
}
