package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIView is a base class for all API views. It provides basic request handling.
type APIView struct{}

// RespondSuccess creates a successful JSON response.
func (v *APIView) RespondSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// RespondCreated creates a successful response for newly created items.
func (v *APIView) RespondCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// RespondError creates an error JSON response with a status code.
func (v *APIView) RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
