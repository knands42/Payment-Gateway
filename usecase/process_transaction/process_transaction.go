package process_transaction

import (
	"log"

	"github.com/caiofernandes00/payment-gateway/adapter/broker"
	"github.com/caiofernandes00/payment-gateway/domain/entity"
	"github.com/caiofernandes00/payment-gateway/domain/repository"
)

type ProcessTransaction struct {
	transactionRepository repository.TransactionRepository
	Producer              broker.ProducerInterface
	Topic                 string
}

func NewProcessTransaction(transactionRepository repository.TransactionRepository, producer broker.ProducerInterface, topic string) *ProcessTransaction {
	return &ProcessTransaction{
		transactionRepository: transactionRepository,
		Producer:              producer,
		Topic:                 topic,
	}
}

func (p *ProcessTransaction) Execute(input TransactionDTOInput) (TransactionDTOOutput, error) {
	log.Printf("Processing transaction %s", input.ID)
	transaction, err := entity.NewTransaction(input.ID, input.AccountId, input.Amount)
	if err != nil {
		return p.handleRejectedTransaction(input, err.Error())
	}
	cc, err := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
	if err != nil {
		return p.handleRejectedTransaction(input, err.Error())
	}

	transaction.SetCreditCard(*cc)

	log.Printf("Transaction %s is valid", input.ID)
	return p.handleApprovedTransaction(input)
}

func (p *ProcessTransaction) handleRejectedTransaction(input TransactionDTOInput, errorMessage string) (TransactionDTOOutput, error) {
	err := p.transactionRepository.Insert(input.ID, input.AccountId, entity.STATUS_REJECTED, errorMessage, input.Amount)
	if err != nil {
		return TransactionDTOOutput{}, err
	}

	output := TransactionDTOOutput{
		ID:           input.ID,
		Status:       entity.STATUS_REJECTED,
		ErrorMessage: errorMessage,
	}

	err = p.publish(output, []byte(input.ID))
	if err != nil {
		return TransactionDTOOutput{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) handleApprovedTransaction(input TransactionDTOInput) (TransactionDTOOutput, error) {
	err := p.transactionRepository.Insert(input.ID, input.AccountId, entity.STATUS_APPROVED, "", input.Amount)
	if err != nil {
		return TransactionDTOOutput{}, err
	}

	output := TransactionDTOOutput{
		ID:           input.ID,
		Status:       entity.STATUS_APPROVED,
		ErrorMessage: "",
	}

	err = p.publish(output, []byte(input.ID))
	if err != nil {
		return TransactionDTOOutput{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) publish(output TransactionDTOOutput, key []byte) error {
	return p.Producer.Publish(output, key, p.Topic)
}
