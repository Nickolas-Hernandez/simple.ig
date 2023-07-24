package main

import (
	"log"

	"simple.ig/database"
	_ "simple.ig/models"

	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	loadEnv()
	loadDatabase()

	router := gin.Default()
	// router.GET("/api/ping", pong)
	// router.POST("/api/user/register", registerUser(db))

	router.Run("localhost:8080")
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func registerUser(c *gin.Context, db *sql.DB) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	// connect to database
	// add new user to database
	// return success or failure

}

func loadDatabase() {
	database.Connect()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
