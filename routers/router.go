package routers

import (
	swaggerFiles "github.com/swaggo/files"
	"goda/config"
	"goda/middlewares"
	"goda/models"
	"goda/serializers"
	"goda/views"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initializes all routes for the application and attaches middleware
func InitRouter() *gin.Engine {
	router := gin.Default()

	// Swagger docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public route: login
	router.POST("/login", views.Login)

	// Protected routes: all routes within this group require authentication
	protected := router.Group("/")              // Create a RouterGroup
	protected.Use(middlewares.AuthMiddleware()) // Add AuthMiddleware to the group

	// Define ViewSet for User and register it under the protected RouterGroup
	userViewSet := &views.ModelViewSet{
		Model:      &models.User{},
		Serializer: &serializers.UserSerializer{},
		DB:         config.DB,
	}
	ViewSetRouter(protected, "users", userViewSet) // Pass the `protected` group instead of the router

	return router
}
