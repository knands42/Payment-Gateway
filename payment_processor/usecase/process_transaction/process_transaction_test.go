package process_transaction

import (
	"context"
	"testing"
	"time"

	tracer_adapter "github.com/caiofernandes00/payment-gateway/adapter/trace"

	mock_broker "github.com/caiofernandes00/payment-gateway/adapter/broker/mock"
	"github.com/caiofernandes00/payment-gateway/domain/entity"
	mock_repository "github.com/caiofernandes00/payment-gateway/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	otel tracer_adapter.TraceClosure = func(ctx context.Context, tracingName string, fn func(context.Context)) context.Context {
		fn(ctx)
		return ctx
	}
	ctx = context.Background()
)

func Test_ProcessTransaction_ExecuteApprovedTransaction(t *testing.T) {
	// Arrange
	input := TransactionDTOInput{
		ID:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "4111111111111111",
		CreditCardName:            "Caio Fernandes",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             "1000",
		Amount:                    200,
	}
	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       entity.STATUS_APPROVED,
		ErrorMessage: "",
	}

	// Act
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	presenterMock := mock_broker.NewMockProducerInterface(ctrl)
	topic := "topic"
	repositoryMock.EXPECT().
		Insert(ctx, input.ID, input.AccountId, expectedOutput.Status, expectedOutput.ErrorMessage, input.Amount).
		Return(nil)

	presenterMock.EXPECT().
		Publish(ctx, expectedOutput, []byte(input.ID), topic).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock, presenterMock, topic, otel)
	output, err := usecase.Execute(ctx, input)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func Test_ProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	// Arrange
	input := TransactionDTOInput{
		ID:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "123",
		CreditCardName:            "Caio Fernandes",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             "1000",
		Amount:                    200,
	}
	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       entity.STATUS_REJECTED,
		ErrorMessage: "creditcard: invalid credit card number",
	}

	// Act
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	presenterMock := mock_broker.NewMockProducerInterface(ctrl)
	topic := "topic"
	repositoryMock.EXPECT().
		Insert(ctx, input.ID, input.AccountId, expectedOutput.Status, expectedOutput.ErrorMessage, input.Amount).
		Return(nil)
	presenterMock.EXPECT().
		Publish(ctx, expectedOutput, []byte(input.ID), topic).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock, presenterMock, topic, otel)

	output, err := usecase.Execute(ctx, input)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func Test_ProcessTransaction_ExecuteRejectedCreditCard(t *testing.T) {
	// Arrange
	input := TransactionDTOInput{
		ID:                        "1",
		AccountId:                 "1",
		CreditCardNumber:          "4111111111111111",
		CreditCardName:            "Caio Fernandes",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             "1000",
		Amount:                    2000,
	}
	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       entity.STATUS_REJECTED,
		ErrorMessage: "transaction: limit exceeded",
	}

	// Act
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	presenterMock := mock_broker.NewMockProducerInterface(ctrl)
	topic := "topic"
	repositoryMock.EXPECT().
		Insert(ctx, input.ID, input.AccountId, expectedOutput.Status, expectedOutput.ErrorMessage, input.Amount).
		Return(nil)
	presenterMock.EXPECT().
		Publish(ctx, expectedOutput, []byte(input.ID), topic).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock, presenterMock, topic, otel)
	output, err := usecase.Execute(ctx, input)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
