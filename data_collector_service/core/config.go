package core

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadConfig() ([]string, string, string) {
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
