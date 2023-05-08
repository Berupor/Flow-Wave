package core

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadMongoConfig() (uri string, dbName string, collectionName string) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dbName = os.Getenv("MONGO_DBNAME")
	collectionName = os.Getenv("MONGO_COLLECTION")

	uri = fmt.Sprintf("mongodb://%s:%s", host, port)

	return uri, dbName, collectionName
}

func LoadRedisConfig() (redisAddr string) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	redisAddr = fmt.Sprintf("%s:%s", host, port)
	return redisAddr
}
