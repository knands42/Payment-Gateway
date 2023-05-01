package process_transaction

import (
	"testing"
	"time"

	mock_broker "github.com/caiofernandes00/payment-gateway/adapter/broker/mock"
	"github.com/caiofernandes00/payment-gateway/domain/entity"
	mock_repository "github.com/caiofernandes00/payment-gateway/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	otel = func(name string, f func()) {
		f()
	}
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
		Insert(input.ID, input.AccountId, expectedOutput.Status, expectedOutput.ErrorMessage, input.Amount).
		Return(nil)

	presenterMock.EXPECT().
		Publish(expectedOutput, []byte(input.ID), topic).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock, presenterMock, topic, otel)
	output, err := usecase.Execute(input)

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
		Insert(input.ID, input.AccountId, expectedOutput.Status, expectedOutput.ErrorMessage, input.Amount).
		Return(nil)
	presenterMock.EXPECT().
		Publish(expectedOutput, []byte(input.ID), topic).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock, presenterMock, topic, otel)

	output, err := usecase.Execute(input)

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
		Insert(input.ID, input.AccountId, expectedOutput.Status, expectedOutput.ErrorMessage, input.Amount).
		Return(nil)
	presenterMock.EXPECT().
		Publish(expectedOutput, []byte(input.ID), topic).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock, presenterMock, topic, otel)
	output, err := usecase.Execute(input)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
