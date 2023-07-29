package middlewares

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DotEnvVariable -> get .env
func DotEnvVariable(key string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
		if key == "PORT" {
			return "8080"
		}
	}
	fmt.Println("Getting value: " + key + ": " + os.Getenv(key))
	return os.Getenv(key)
}
