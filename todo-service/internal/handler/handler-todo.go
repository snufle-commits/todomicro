package handler

import (
	"net/http"
	"strconv"
	"todoservice/internal/model"
	"todoservice/internal/service"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{service: *s}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var todo model.ToDo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	userIDAny, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userID := userIDAny.(uint)

	todos, err := h.service.GetAll(c, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	userIDAny, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userID, ok := userIDAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	todo, err := h.service.GetByID(c, userID, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Complete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	userIDAny, exists := c.Get("UserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userID, ok := userIDAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	if err := h.service.Complete(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "completed"})
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
