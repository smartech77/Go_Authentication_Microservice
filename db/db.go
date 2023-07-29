package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	mongo "github.com/helios/go-sdk/proxy-libs/heliosmongo"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var client *mongo.Client

func Dbconnect() *mongo.Client {

	clientOptions := options.Client().ApplyURI(DotEnvVariable("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")
	return client
}
