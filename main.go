package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
