package process_transaction

import (
	"github.com/caiofernandes00/payment-gateway/domain/entity"
	"github.com/caiofernandes00/payment-gateway/domain/repository"
)

type ProcessTransaction struct {
	transactionRepository repository.TransactionRepository
}

func NewProcessTransaction(transactionRepository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{transactionRepository: transactionRepository}
}

func (p *ProcessTransaction) Execute(input TransactionDTOInput) (TransactionDTOOutput, error) {
	transaction, err := entity.NewTransaction(input.ID, input.AccountId, input.Amount)
	if err != nil {
		return p.handleRejectedTransaction(input, err.Error())
	}
	cc, err := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
	if err != nil {
		return p.handleRejectedTransaction(input, err.Error())
	}

	transaction.SetCreditCard(*cc)

	return p.handleApprovedTransaction(input)
}

func (p *ProcessTransaction) handleRejectedTransaction(input TransactionDTOInput, errorMessage string) (TransactionDTOOutput, error) {
	err := p.transactionRepository.Insert(input.ID, input.AccountId, entity.STATUS_REJECTED, errorMessage, input.Amount)
	if err != nil {
		return TransactionDTOOutput{}, err
	}

	return TransactionDTOOutput{
		ID:           input.ID,
		Status:       entity.STATUS_REJECTED,
		ErrorMessage: errorMessage,
	}, nil
}

func (p *ProcessTransaction) handleApprovedTransaction(input TransactionDTOInput) (TransactionDTOOutput, error) {
	err := p.transactionRepository.Insert(input.ID, input.AccountId, entity.STATUS_APPROVED, "", input.Amount)
	if err != nil {
		return TransactionDTOOutput{}, err
	}

	return TransactionDTOOutput{
		ID:           input.ID,
		Status:       entity.STATUS_APPROVED,
		ErrorMessage: "",
	}, nil
}
