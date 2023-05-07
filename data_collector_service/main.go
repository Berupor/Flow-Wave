package main

import (
	"fmt"

	"data-collector/handlers"
	"data-collector/kafka"
	"data-collector/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	brokers := []string{"localhost:9092"}
	topic := "reviews"

	messageProducer, err := kafka.NewKafkaProducer(
		brokers, topic,
	)

	if err != nil {
		fmt.Printf("failed to create Kafka producer: %s\n", err)
	}

	eventService := services.NewEventService(
		*messageProducer,
	)

	r.POST("/review", handlers.CreateReview(*eventService))
	r.Run()
}
