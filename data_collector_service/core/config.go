package core

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func KafkaConfig() ([]string, string, string) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error loading file .env: %v", err)
	}

	brokersString := os.Getenv("KAFKA_BROKERS")
	brokers := strings.Split(brokersString, ",")
	topic := os.Getenv("KAFKA_TOPIC")
	jwtSecret := os.Getenv("JWT_SECRET")

	return brokers, topic, jwtSecret
}

func ApiConfig() string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error loading file .env: %v", err)
	}

	host := os.Getenv("GO_DATA_COLLECTOR_HOST")
	port := os.Getenv("GO_DATA_COLLECTOR_PORT")

	return host + ":" + port
}
