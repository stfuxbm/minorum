package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}
}

func GetMongoURI() string {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("mongo_uri not found")
	}
	return uri
}

func GetMongoDBName() string {
	db := os.Getenv("MONGO_DB")
	if db == "" {
		log.Fatal("mongo_db not found")
	}
	return db
}
