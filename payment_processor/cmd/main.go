package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/caiofernandes00/payment-gateway/adapter/broker/kafka"
	"github.com/caiofernandes00/payment-gateway/adapter/factory"
	"github.com/caiofernandes00/payment-gateway/adapter/presenter/transaction"
	repository_adapter "github.com/caiofernandes00/payment-gateway/adapter/repository"
	"github.com/caiofernandes00/payment-gateway/domain/repository"
	"github.com/caiofernandes00/payment-gateway/usecase/process_transaction"
	"github.com/caiofernandes00/payment-gateway/util"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
)

var (
	config            *util.Config
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
		log.Println("Message received" + string(msg.Value))
		var input process_transaction.TransactionDTOInput
		json.Unmarshal(msg.Value, &input)
		log.Println("Message unmarshalled")
		log.Println(input)
		usecase.Execute(input)
	}
}

func loadEnv() {
	config = util.NewConfig()
	path, err := getRootFile()

	if err != nil {
		log.Println("No env file provided, using only env variables")
	}

	config.LoadEnv(path)
}

func getRootFile() (ex string, err error) {
	ex, _ = os.Getwd()
	_, err = os.Stat(filepath.Join(ex, "app.env"))

	if err != nil {
		ex = filepath.Join(ex, "../")
		_, err = os.Stat(filepath.Join(ex, "app.env"))

		if err != nil {
			log.Println("No env file provided, using only env variables")
		}
	}

	return
}

func initializeDb() {
	var err error
	db, err = sql.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		panic(err)
	}
	initializeMigration(db)
}

func initializeMigration(db *sql.DB) {
	ex, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(ex, "cmd") {
		ex = filepath.Join(ex, "../")
	}
	migrationsDir := os.DirFS(filepath.Join(ex, "/adapter/repository/migrations"))
	repository_adapter.Up(db, migrationsDir)
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
