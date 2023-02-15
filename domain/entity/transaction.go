package entity

import (
	"errors"

	NotificationPackage "github.com/caiofernandes00/payment-gateway/domain/notification"
)

var TRANSACTION_CONTEXT string = "transaction"

type Transaction struct {
	ID           string
	AccountId    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
	notification NotificationPackage.Notification
}

func NewTransaction(id, accountId string, amount float64) (*Transaction, error) {
	t := &Transaction{
		ID:        id,
		AccountId: accountId,
		Amount:    amount,
	}

	if err := t.IsValid(); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Transaction) IsValid() error {
	t.validateAmount()

	if t.notification.HasErrors() {
		return errors.New(t.notification.Messages(TRANSACTION_CONTEXT))
	}

	return nil
}

func (t *Transaction) SetCreditCard(cc CreditCard) {
	t.CreditCard = cc
}

func (t *Transaction) validateAmount() {
	if t.Amount <= 0 {
		t.notification.AddError("the amount must be greater than 0", TRANSACTION_CONTEXT)
	}

	if t.Amount > 1000 {
		t.notification.AddError("limit exceeded", TRANSACTION_CONTEXT)
	}
}
