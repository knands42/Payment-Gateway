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
	log.Println("Starting consumer...")
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		log.Println("Failed to create consumer: " + err.Error())
		return err
	}

	// TODO: Add rebalance callback
	log.Println("Starting to consume messages...")
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		log.Println("Failed to subscribe to topic: " + err.Error())
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
