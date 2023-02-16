package main

import (
	"database/sql"
	"encoding/json"

	"github.com/caiofernandes00/payment-gateway/adapter/broker/kafka"
	"github.com/caiofernandes00/payment-gateway/adapter/presenter/transaction"
	"github.com/caiofernandes00/payment-gateway/domain/repository"
	"github.com/caiofernandes00/payment-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	config            *Config
	db                *sql.DB
	repo              repository.TransactionRepository
	usecase           *process_transaction.ProcessTransaction
	kafkaPresenter    *transaction.TransactionKafkaPresenter
	kafkaProducer     *kafka.Producer
	kafkaConsumer     *kafka.Consumer
	kafkaConsumerChan chan *ckafka.Message = make(chan *ckafka.Message)
)

func init() {
	loadEnv()
	initializeDb()
	initializeKafka()
	initializeRepo()
	initializeUsecase()
}

func main() {
	go kafkaConsumer.Consume(kafkaConsumerChan)
	for msg := range kafkaConsumerChan {
		var input process_transaction.TransactionDTOInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
