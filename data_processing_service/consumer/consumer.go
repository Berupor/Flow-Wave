package consumer

import (
	"context"

	"data-processing/models"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

type Handler interface {
	HandleMessage(event models.Review) error
}

type KafkaConsumer struct {
	ConsumerGroup sarama.ConsumerGroup
	Handler       Handler
}

func NewKafkaConsumer(brokers []string, groupID string, topics []string, handler Handler) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0
	config.Consumer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	consumer := &KafkaConsumer{
		ConsumerGroup: consumerGroup,
		Handler:       handler,
	}

	go func() {
		for {
			err := consumerGroup.Consume(context.Background(), topics, consumer)
			if err != nil {
				log.Printf("Error from consumer: %v", err)
			}
		}
	}()

	return consumer, nil
}

func (c *KafkaConsumer) Setup(sarama.ConsumerGroupSession) error {
	log.Println("Consumer setup")
	return nil
}

func (c *KafkaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("Consumer cleanup")
	return nil
}

func (c *KafkaConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Message received: %s", string(msg.Value))

		var event models.Review
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("Error decoding JSON message: %v", err)
			continue
		}

		err = c.Handler.HandleMessage(event)
		if err != nil {
			log.Printf("Error processing message: %v", err)
		} else {
			sess.MarkMessage(msg, "")
		}
	}

	return nil
}

func (c *KafkaConsumer) Close() error {
	return c.ConsumerGroup.Close()
}
