package router

import (
	"todoservice/internal/handler"
	"todoservice/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.TodoHandler) *gin.Engine {
	r := gin.Default()
	api := r.Group("/todos")
	api.Use(middleware.RequireAuth)
	{
		api.POST("/", handler.Create)
		api.GET("/", handler.GetAll)
		api.GET("/:id", handler.GetByID)
		api.PATCH("/id", handler.Complete)
		api.DELETE("/:id", handler.Delete)
	}
	return r

}
