package handler

import (
	"authservice/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: *service}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse body"})
		return
	}

	if err := h.service.SignUp(body.Email, body.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to sign up"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully signed up",
	})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse body"})
		return
	}

	token, err := h.service.SignIn(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to sign in"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
