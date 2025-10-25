package router

import (
	"authservice/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()
	api := r.Group("/")
	{
		api.POST("/login", handler.SignIn)
		api.POST("/register", handler.SignUp)
	}
	return r

}
