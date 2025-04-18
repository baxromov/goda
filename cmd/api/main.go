package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"

	"goda/api/routes"
	"goda/configs"
	"goda/pkg/handlers"
	"goda/pkg/middlewares"
	"goda/pkg/repositories"
	"goda/pkg/services"
)

func main() {
	// Load Configurations
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Initialize Database
	db, err := configs.InitDatabase(config)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	// Dependency Injection
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// Router
	router := gin.Default()

	// Middleware (dynamic based on config)
	if config.Middleware.Logging {
		router.Use(gin.Logger())
	}
	if config.Middleware.Recovery {
		router.Use(gin.Recovery())
	}
	if config.Middleware.Auth {
		router.Use(middlewares.AuthMiddleware())
	}

	// Routes
	routes.RegisterUserRoutes(router, userHandler)

	// Run Server
	routerAddress := fmt.Sprintf(":%d", config.Server.Port)
	log.Printf("Starting server on %s", routerAddress)
	router.Run(routerAddress)
}
