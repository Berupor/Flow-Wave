package consumer

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
)

type Handler interface {
	HandleMessage(message string) error
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

		err := c.Handler.HandleMessage(string(msg.Value))
		if err != nil {
			log.Printf("Error processing message: %v", err)
		} else {
			sess.MarkMessage(msg, "")
		}
	}
	return nil
}
