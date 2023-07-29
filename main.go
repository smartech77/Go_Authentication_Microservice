package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DotEnvVariable(key string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println("Getting value: " + key + ": " + os.Getenv(key))
	return os.Getenv(key)
}

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"username" bson:"username" `
	Password  string             `json:"password" bson:"password"`
	Email     string             `json:"email" bson:"email"`
	Language  string             `json:"language" bson:"language"`
	Phone     string             `json:"phone" bson:"phone"`
	Firstname string             `json:"firstname" bson:"firstname" validate:"required,alpha"`
	Lastname  string             `json:"lastname" bson:"lastname" validate:"required,alpha"`
	Photo     string             `json:"photo" bson:"photo"`
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
