package middlewares

import (
	"goda/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies the JWT token and sets user info in the context.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			c.Abort()
			return
		}

		// Split "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header format"})
			c.Abort()
			return
		}

		// Parse and validate the token
		claims, err := auth.ParseJWT(tokenParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Set user info in the context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}

// IsAdmin checks if the user is an admin.
func IsAdmin(c *gin.Context) bool {
	// Replace this with a real admin check (e.g., from the database)
	return c.GetString("email") == "admin@example.com"
}

// RequireAdmin middleware ensures only admins can access a route.
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsAdmin(c) {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
