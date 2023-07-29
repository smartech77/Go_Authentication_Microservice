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
	// we should add color package to use for logger here.
	if err != nil {
		log.Fatalf("Error loading .env file")
		if os.Getenv(key) == "" && key == "PORT" {
			return "8080"
		}
	}
	fmt.Println("Getting value: " + key + ": " + os.Getenv(key))

	
	return os.Getenv(key)
}
