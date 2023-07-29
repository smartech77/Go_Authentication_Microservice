package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	port := middlewares.DotEnvVariable("PORT")
	fmt.Println("Port number is: " + port)
	color.Cyan("🌏 Server running on localhost:" + port)

	router := gin.Default()

	router := gin.Default()

	// router.Use(customLogger())
	routes.Routes(router)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
