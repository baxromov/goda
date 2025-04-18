package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware ensures the user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		// Proceed if authenticated (validate token logic here)
		c.Next()
	}
}
