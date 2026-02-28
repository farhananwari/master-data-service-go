package main

import (
	"fmt"
	"log"
	"os"

	"master-data-service-go/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "up":
		if err := database.RunMigrations(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migration completed successfully!")

	case "down":
		if err := database.RollbackMigration(); err != nil {
			log.Fatalf("Rollback failed: %v", err)
		}
		log.Println("Rollback completed successfully!")

	case "version":
		version, dirty, err := database.MigrationVersion()
		if err != nil {
			log.Fatalf("Failed to get migration version: %v", err)
		}
		fmt.Printf("Current migration version: %d (dirty: %v)\n", version, dirty)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: go run cmd/migrate/main.go [command]")
	fmt.Println("\nCommands:")
	fmt.Println("  up       - Run all pending migrations")
	fmt.Println("  down     - Rollback the last migration")
	fmt.Println("  version  - Show current migration version")
}
