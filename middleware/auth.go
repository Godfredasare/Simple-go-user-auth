package middleware

import (
	"net/http"
	"simple/user/auth/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token required"})
		c.Abort()
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	// Set the userId to the context
	c.Set("userId", userId)
	c.Next()

}
