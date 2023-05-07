package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

type Producer interface {
	Close() error
	Produce(message interface{}) error
}

type KafkaProducer struct {
	Producer sarama.SyncProducer
	Topic    string
}

func NewKafkaProducer(brokers []string, topic string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		Producer: producer,
		Topic:    topic,
	}, nil
}

func (p *KafkaProducer) Produce(message interface{}) error {
	messageBytes, err := json.Marshal(message)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.StringEncoder(messageBytes),
	}

	partition, offset, err := p.Producer.SendMessage(msg)

	if err != nil {
		return err
	}

	fmt.Printf("Message stored in topic(%s)/partition(%d)/offset(%d)\n", p.Topic, partition, offset)
	return nil
}
