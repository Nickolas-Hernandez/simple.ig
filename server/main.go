package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", sayHello)

	router.Run("localhost:8080")
}

func sayHello(c *gin.Context) {
	hello := "Hello World"
	c.IndentedJSON(http.StatusOK, hello)
}
