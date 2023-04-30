package main

import (
	"context"
	"database/sql"
	"encoding/json"
	tracer_adapter "github.com/caiofernandes00/payment-gateway/adapter/trace"
	"github.com/caiofernandes00/payment-gateway/adapter/trace/exporter"
	"go.opentelemetry.io/otel/trace"
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
	"go.opentelemetry.io/otel/baggage"
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
	otel              trace.Tracer
)

func init() {
	loadEnv()
	initializeTrace()
	initializeDb()
	initializeKafka()
	initializeRepo()
	initializeUseCase()
}

func main() {
	ctx := baggage.ContextWithoutBaggage(context.Background())

	go func() {
		tracer_adapter.TraceFn(otel, ctx, "kafka-consumer-listener", func() {
			err := kafkaConsumer.Consume(kafkaConsumerChan)
			if err != nil {
				log.Println("Error to consume kafka message" + err.Error())
			}
		})
	}()

	for msg := range kafkaConsumerChan {
		var err error
		var input process_transaction.TransactionDTOInput

		tracer_adapter.TraceFn(otel, ctx, "kafka-consumer-reader", func() {
			log.Println("Message received" + string(msg.Value))
			err = json.Unmarshal(msg.Value, &input)
		})
		if err != nil {
			log.Println("Error to unmarshal message" + err.Error())
			continue
		}

		tracer_adapter.TraceFn(otel, ctx, "usecase-process-transaction", func() {
			log.Println("Message unmarshalled")
			log.Println(input)
			_, err = usecase.Execute(input)
		})
		if err != nil {
			log.Println("Error to process transaction" + err.Error())
			continue
		}
	}
}

func loadEnv() {
	config = util.NewConfig()
	config.LoadEnv(config.Profile)
}

func initializeTrace() {
	zipkin := exporter.NewZipkinExporter(config.ExporterEndpoint)
	otel = tracer_adapter.NewOpenTelemetry(zipkin.GetExporter()).GetTracer()
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

func initializeUseCase() {
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
