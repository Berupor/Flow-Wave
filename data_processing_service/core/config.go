package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadKafkaConfig() ([]string, []string, string) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error loading file .env: %v", err)
	}

	brokersString := os.Getenv("KAFKA_BROKERS")
	brokers := strings.Split(brokersString, ",")

	topicsString := os.Getenv("KAFKA_TOPIC")
	topics := strings.Split(topicsString, ",")

	groupID := os.Getenv("KAFKA_GROUP_ID")

	return brokers, topics, groupID
}

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
