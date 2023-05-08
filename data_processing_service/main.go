package main

import (
	"data-processing/storage"
	"log"
	"os"
	"os/signal"
	"syscall"

	"data-processing/consumer"
	"data-processing/core"
	"data-processing/processor"
)

func main() {
	brokers, topics, groupID := core.LoadKafkaConfig()
	connectionUri, dbName, collection := core.LoadMongoConfig()

	mongo, err := storage.NewMongoStorage(connectionUri, dbName, collection)
	if err != nil {
		log.Fatalf("error creating mongo connection: %v", err)
	}

	handler := &processor.ReviewAnalyzerHandler{Storage: mongo}
	kafkaConsumer, err := consumer.NewKafkaConsumer(brokers, groupID, topics, handler)
	if err != nil {
		log.Fatalf("error creating Kafka consumer: %v", err)
	}

	// Waiting for signals to terminate the application
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// block the execution of main until we receive a stop signal
	<-signals

	log.Println("shutting down the service...")

	if err := kafkaConsumer.Close(); err != nil {
		log.Printf("error closing Kafka consumer: %v", err)
	}
}
