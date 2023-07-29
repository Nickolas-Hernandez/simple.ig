package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// set up env variables
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}

	// Create db connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	defer conn.Close(context.Background())

	// Create router
	app := fiber.New()

	// Create routes
	app.Get("/ping", pingHandler)

	// Start server
	app.Listen(":3000")

	fmt.Println("Hello World!")
}

func pingHandler(c *fiber.Ctx) error {
	return c.SendString("pong")
}
