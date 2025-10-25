package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil || tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte("gfdsgsfgfsg"), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Println("JWT parse error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID, ok := claims["id"].(string)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("UserID", userID)
	c.Next()
}
