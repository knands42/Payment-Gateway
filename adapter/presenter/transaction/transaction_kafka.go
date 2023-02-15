package transaction

import (
	"encoding/json"
	"errors"

	"github.com/caiofernandes00/payment-gateway/usecase/process_transaction"
)

type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewKafkaPresenter(kafkaPresenter KafkaPresenter) *KafkaPresenter {
	return &KafkaPresenter{}
}

func (k *KafkaPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(k)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (k *KafkaPresenter) Bind(input interface{}) error {
	inputCast, ok := input.(process_transaction.TransactionDTOOutput)
	if !ok {
		return errors.New("invalid input")
	}

	k.ID = inputCast.ID
	k.Status = inputCast.Status
	k.ErrorMessage = inputCast.ErrorMessage

	return nil
}
