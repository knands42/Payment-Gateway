package transaction

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/caiofernandes00/payment-gateway/usecase/process_transaction"
)

type TransactionKafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewTransactionKafkaPresenter() *TransactionKafkaPresenter {
	return &TransactionKafkaPresenter{}
}

func (k *TransactionKafkaPresenter) Show() ([]byte, error) {
	log.Printf("Showing transaction: %v", k)
	j, err := json.Marshal(k)
	if err != nil {
		log.Printf("Failed to show transaction: %v", err)
		return nil, err
	}

	log.Printf("Showed transaction: %v", k)
	return j, nil
}

func (k *TransactionKafkaPresenter) Bind(input interface{}) error {
	log.Printf("Binding input: %v", input)
	inputCast, ok := input.(process_transaction.TransactionDTOOutput)
	if !ok {
		log.Printf("Failed to bind: %v", input)
		return errors.New("invalid input")
	}

	k.ID = inputCast.ID
	k.Status = inputCast.Status
	k.ErrorMessage = inputCast.ErrorMessage

	log.Printf("Binded input: %v", input)
	return nil
}
