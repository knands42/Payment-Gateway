package util

import (
	"fmt"
	"os"

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
		KafkaBootstrapServers: "payment_base_kafka:9094",
		KafkaProducerTopic:    "transactions_result",
		KafkaConsumerTopic:    "transactions",
		KafkaConsumerClientId: "payment_processor",
		KafkaConsumerGroupId:  "payment_processor",
	}
}

func (c *Config) LoadEnv() {
	viper.AutomaticEnv()

	err := viper.Unmarshal(&c)

	v := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if v != "" {
		c.KafkaBootstrapServers = v
	}

	if err != nil {
		fmt.Printf("Error unmarshalling config, %s", err)
	}
}
