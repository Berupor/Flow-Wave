package services

import (
	"data-collector/kafka"
	"log"
)

type EventService struct {
	Producer kafka.KafkaProducer
}

func NewEventService(producer kafka.KafkaProducer) *EventService {
	return &EventService{
		Producer: producer,
	}
}

func (s *EventService) Produce(message interface{}) error {

	if err := s.Producer.Produce(message); err != nil {
		log.Printf("failed to produce message: %s", err)
	}
	return nil
}
