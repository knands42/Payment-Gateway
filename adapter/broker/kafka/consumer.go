package kafka

import (
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewKafkaConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		return err
	}

	// TODO: Add rebalance callback
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		return err
	}

	for {
		log.Println("Waiting for messages...")
		msg, err := consumer.ReadMessage(-1)
		log.Println("Message received" + string(msg.Value))
		if err == nil {
			msgChan <- msg
		}
	}
}
