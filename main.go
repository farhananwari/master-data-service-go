package main

import (
	"log"
	"os"

	"master-data-service-go/database"
	"master-data-service-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables first
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	}

	// Connect to database FIRST
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Get DB instance AFTER connection is established
	db := database.GetDB()

	// Validate DB is not nil
	if db == nil {
		log.Fatal("Database connection is nil after Connect()")
	}

	log.Println("Starting master-data-service-go...")

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, db)

	// Get server configuration
	serverPort := getEnv("SERVER_PORT", "8080")
	serverHost := getEnv("SERVER_HOST", "localhost")

	// Start server
	address := serverHost + ":" + serverPort
	log.Printf("Server is running on http://%s\n", address)

	if err := router.Run(":" + serverPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Helper function to get environment variables with default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
