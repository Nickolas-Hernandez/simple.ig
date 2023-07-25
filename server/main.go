package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// set up env variables
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}

	// dbHost := os.Getenv("DB_HOST")

	// Create db connection
	// Create router
	// Create routes
	// Start server

	fmt.Println("Hello World!")
}
