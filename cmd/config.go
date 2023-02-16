package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DriverName            string `mapstructure:"DRIVER_NAME"`
	DataSourceName        string `mapstructure:"DATA_SOURCE_NAME"`
	KafkaBootstrapServers string `mapstructure:"KAFKA_BOOTSTRAP_SERVERS"`
	KafkaProducerTopic    string `mapstructure:"KAFKA_PRODUCER_TOPIC"`
	KafkaConsumerTopic    string `mapstructure:"KAFKA_CONSUMER_TOPIC"`
	KafkaConsumerClientId string `mapstructure:"KAFKA_CONSUMER_CLIENT_ID"`
	KafkaConsumerGroupId  string `mapstructure:"KAFKA_CONSUMER_GROUP_ID"`
}

func NewConfig() *Config {
	return &Config{
		DriverName:            "sqlite3",
		DataSourceName:        "transaction.db",
		KafkaBootstrapServers: "localhost:9092",
		KafkaProducerTopic:    "transactions_result",
		KafkaConsumerTopic:    "transactions",
		KafkaConsumerClientId: "goapp",
		KafkaConsumerGroupId:  "goapp",
	}
}

func (c *Config) LoadEnv() (config Config, err error) {
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	fmt.Printf("Error unmarshalling config, %s", err)
	return
}
