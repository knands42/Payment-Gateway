package kafka

import (
	"context"
	"testing"

	tracer_adapter "github.com/caiofernandes00/payment-gateway/adapter/trace"
	"github.com/golang/mock/gomock"

	"github.com/caiofernandes00/payment-gateway/adapter/presenter/transaction"
	"github.com/caiofernandes00/payment-gateway/domain/entity"
	"github.com/caiofernandes00/payment-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

var (
	otel tracer_adapter.TraceClosure = func(ctx context.Context, tracingName string, fn func(context.Context)) context.Context {
		fn(ctx)
		return ctx
	}
	ctx = context.Background()
)

func Test_ProducerPublish(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expectedOutput := process_transaction.TransactionDTOOutput{
		ID:           "1",
		Status:       entity.STATUS_REJECTED,
		ErrorMessage: "limit exceeded",
	}

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	// Act
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter(), otel)
	err := producer.Publish(ctx, expectedOutput, []byte("1"), "test")

	// Assert
	assert.Nil(t, err)
}
