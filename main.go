package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func DotEnvVariable(key string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println("Getting value: " + key + ": " + os.Getenv(key))
	return os.Getenv(key)
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the homepage
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the homepage!")
	})

	// Define a route for a user profile page
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "This is the profile page for user %s", id)
	})

	// Run the router on port 8080
	router.Run(":8080")
}
