package main

import (
	"data-collector/core"
	"data-collector/kafka"
	"data-collector/services"
	"log"
)

func main() {
	brokers, topic, jwtSecret := core.KafkaConfig()
	apiAddr := core.ApiConfig()

	messageProducer, err := kafka.NewKafkaProducer(brokers, topic)
	if err != nil {
		log.Fatalf("failed to create Kafka producer: %s", err)
	}

	eventService := services.NewEventService(*messageProducer)

	router := core.SetupRouter(eventService, jwtSecret)
	router.Run(apiAddr)
}
