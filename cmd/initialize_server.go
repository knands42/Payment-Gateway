package main

import (
	"database/sql"

	"github.com/caiofernandes00/payment-gateway/adapter/broker/kafka"
	"github.com/caiofernandes00/payment-gateway/adapter/factory"
	"github.com/caiofernandes00/payment-gateway/adapter/presenter/transaction"
	"github.com/caiofernandes00/payment-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
)

func loadEnv() {
	config = NewConfig()
	config.LoadEnv()
}

func initializeDb() {
	var err error
	db, err = sql.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		panic(err)
	}
}

func initializeRepo() {
	repo = factory.
		NewRepositoryDatabaseFactory(db).
		CreateTransactionRepository()
}

func initializeUsecase() {
	usecase = process_transaction.NewProcessTransaction(repo, kafkaProducer, config.KafkaProducerTopic)
}

func initializeKafka() {
	kafkaPresenter = transaction.NewTransactionKafkaPresenter()
	initializeKafkaConsumer()
	initializeKafkaProducer(kafkaPresenter)
}

func initializeKafkaConsumer() {
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": config.KafkaBootstrapServers,
		"client.id":         config.KafkaConsumerClientId,
		"group.id":          config.KafkaConsumerGroupId,
	}
	topics := []string{config.KafkaConsumerTopic}
	kafkaConsumer = kafka.NewKafkaConsumer(configMapConsumer, topics)
}

func initializeKafkaProducer(presenter *transaction.TransactionKafkaPresenter) {
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": config.KafkaBootstrapServers,
	}
	kafkaProducer = kafka.NewKafkaProducer(configMapProducer, presenter)
}
