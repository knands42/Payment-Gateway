package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Profile               string `mapstructure:"PROFILE"`
	DriverName            string `mapstructure:"DRIVER_NAME"`
	DataSourceName        string `mapstructure:"DATA_SOURCE_NAME"`
	KafkaBootstrapServers string `mapstructure:"KAFKA_BOOTSTRAP_SERVERS"`
	KafkaProducerTopic    string `mapstructure:"KAFKA_PRODUCER_TOPIC"`
	KafkaConsumerTopic    string `mapstructure:"KAFKA_CONSUMER_TOPIC"`
	KafkaConsumerClientId string `mapstructure:"KAFKA_CONSUMER_CLIENT_ID"`
	KafkaConsumerGroupId  string `mapstructure:"KAFKA_CONSUMER_GROUP_ID"`
	ExporterEndpoint      string `mapstructure:"EXPORTER_ENDPOINT"`
	NewRelicConfigAppName string `mapstructure:"NEW_RELIC_APP_NAME"`
	NewRelicConfigLicense string `mapstructure:"NEW_RELIC_LICENSE"`
}

func NewConfig() *Config {
	return &Config{
		Profile:               "local",
		DriverName:            "sqlite3",
		DataSourceName:        "transaction.db",
		KafkaBootstrapServers: "payment_base_kafka:9094",
		KafkaProducerTopic:    "transactions_result",
		KafkaConsumerTopic:    "transactions",
		KafkaConsumerClientId: "payment_processor",
		KafkaConsumerGroupId:  "payment_processor",
		ExporterEndpoint:      "http://localhost:9411/api/v2/spans",
		NewRelicConfigAppName: "Payment-Processor",
	}
}

func (c *Config) LoadEnv(env string) {
	path, _ := getRootFile(env)

	viper.AddConfigPath(path)
	viper.SetConfigName("app." + env)
	viper.SetConfigType("env")

	err := viper.Unmarshal(&c)

	viper.AutomaticEnv()

	v := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if v != "" {
		c.KafkaBootstrapServers = v
	}
	v = os.Getenv("EXPORTER_ENDPOINT")
	if v != "" {
		c.ExporterEndpoint = v
	}

	if err != nil {
		fmt.Printf("Error unmarshalling config, %s", err)
	}
}

func getRootFile(env string) (ex string, err error) {
	ex, _ = os.Getwd()
	_, err = os.Stat(filepath.Join(ex, "app."+env+".env"))

	if err != nil {
		for i := 0; i < 5; i++ {
			ex = filepath.Join(ex, "../")
			_, err = os.Stat(filepath.Join(ex, "app."+env+".env"))

			if err == nil {
				break
			}
		}
		if err != nil {
			log.Println("No env file provided, using only env variables")
		}
	}

	return
}
