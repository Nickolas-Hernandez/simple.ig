package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// var db *pgx.Conn

type App struct {
	DB *pgx.Conn
}

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

	db := App{
		DB: conn,
	}

	// Create router
	startServer(db)
}

func startServer(db App) {
	app := fiber.New()

	app.Get("/ping", pingHandler)
	app.Post("/api/auth/register", db.loginHandler)

	app.Listen(":3000")
}

func (db *App) loginHandler(c *fiber.Ctx) error {
	fmt.Println("hello")
	var user User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	fmt.Println(user.Username)
	fmt.Println(user.Password)
	fmt.Println(user.Email)

	if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and password are required",
		})
	}

	// query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	// _, err = db.Exec(context.Background(), query, username, hashedPassword)
	// if err != nil {
	// 	return err
	// }

	// body := string(c.Body())
	// fmt.Println("Request body: ", body)
	// c.Status(fiber.StatusOK)
	return c.SendString("Body has been logged")

}

func pingHandler(c *fiber.Ctx) error {
	return c.SendString("pong")
}
