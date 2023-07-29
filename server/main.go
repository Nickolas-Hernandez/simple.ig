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
	app.Get("/api/ping", pingHandler)
	app.Post("/api/auth/register", loginHandler)

	// Start server
	app.Listen(":3000")

	fmt.Println("Hello World!")
}

func loginHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
	body := string(c.Body())
	fmt.Println("Request body: ", body)
	c.Status(fiber.StatusOK)
	return c.SendString("Body has been logged")
}

func pingHandler(c *fiber.Ctx) error {
	return c.SendString("pong")
}
