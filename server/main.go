package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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
	app.Post("/api/auth/register", db.registerUser)

	app.Listen(":3000")
}

func (db *App) registerUser(c *fiber.Ctx) error {
	var user User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	username, password, email := user.Username, user.Password, user.Email

	if strings.TrimSpace(username) == "" ||
		strings.TrimSpace(password) == "" ||
		strings.TrimSpace(email) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and password are required",
		})
	}

	hashedPassword, err := HashPassword(password)

	// fmt.Println("hashedPassword: ", hashedPassword)

	// fmt.Print(db.DB.Query(context.Background(), "select current_database()"))

	query := `insert into public.users("username", "email", "password") values($1, $2, $3)`
	_, err = db.DB.Exec(context.Background(), query, username, email, hashedPassword)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  username,
		"email": email,
	})

	key := os.Getenv("JWT_SECRET")

	keytype := reflect.TypeOf(key)

	fmt.Println("keytype: ", keytype)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("error: ", err)

		return err
	}

	c.Status(fiber.StatusOK)

	fmt.Print("status set")

	return c.JSON(fiber.Map{
		"token":  signedToken,
		username: username,
		email:    email,
	})

}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func pingHandler(c *fiber.Ctx) error {
	return c.SendString("pong")
}
