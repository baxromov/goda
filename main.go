package main

import (
	"goda/config"
	"goda/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	// Initialize the database
	config.ConnectDatabase()

	// Initialize the router
	router := routers.InitRouter()

	// Start the server on the specified port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if PORT is not set
	}

	log.Printf("Server is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
