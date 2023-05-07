package main

import (
	"fmt"

	"data-collector/handlers"
	"data-collector/kafka"
	"data-collector/middlewares"
	"data-collector/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

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

	secured := router.Group("api")
	secured.Use(middlewares.JWTAuthMiddleware())
	{
		secured.POST("/review", handlers.CreateReview(*eventService))
	}

	// r.POST("/review", handlers.CreateReview(*eventService))
	router.Run()
}
