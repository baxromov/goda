package routes

import (
	"github.com/gin-gonic/gin"
	"goda/pkg/handlers"
)

func RegisterUserRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", handler.Register)
	}
}
