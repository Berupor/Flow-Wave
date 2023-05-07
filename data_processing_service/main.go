package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"data-processing/consumer"
	"data-processing/core"
	"data-processing/processor"
)

func main() {
	brokers, topics, groupID := core.LoadConfig()

	handler := &processor.SimpleHandler{}
	kafkaConsumer, err := consumer.NewKafkaConsumer(brokers, groupID, topics, handler)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}

	// Ждем сигналов завершения работы приложения
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Блокируем выполнение main, пока не получим сигнал остановки
	<-signals

	log.Println("Shutting down the service...")

	if err := kafkaConsumer.Close(); err != nil {
		log.Printf("Error closing Kafka consumer: %v", err)
	}
}
