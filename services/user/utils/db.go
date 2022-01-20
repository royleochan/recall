package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDb(ctx context.Context) (*mongo.Client, error, string) {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	// Connect to DB
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	MONGO_CONN_STRING := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.karg0.mongodb.net/%s?retryWrites=true&w=majority", DB_USER, DB_PASSWORD, DB_NAME)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_CONN_STRING))
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	return client, err, DB_NAME
}
