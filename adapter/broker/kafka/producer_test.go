package kafka

import (
	"testing"

	"github.com/caiofernandes00/payment-gateway/adapter/presenter/transaction"
	"github.com/caiofernandes00/payment-gateway/domain/entity"
	"github.com/caiofernandes00/payment-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func Test_ProducerPublish(t *testing.T) {
	// Arrange
	expectedOutput := process_transaction.TransactionDTOOutput{
		ID:           "1",
		Status:       entity.STATUS_REJECTED,
		ErrorMessage: "limit exceeded",
	}

	// outputJson, _ := json.Marshal(expectedOutput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	// Act
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")

	// Assert
	assert.Nil(t, err)
}
