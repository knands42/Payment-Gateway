package entity

import "errors"

type Transaction struct {
	ID           string
	AccountId    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
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
	if !t.validateID() {
		return errors.New("invalid transaction id")
	}

	if err := t.validateAmount(); err != nil {
		return err
	}

	return nil
}

func (t *Transaction) SetCreditCard(cc CreditCard) {
	t.CreditCard = cc
}

func (t *Transaction) validateID() bool {
	return t.ID != ""
}

func (t *Transaction) validateAmount() error {
	if t.Amount < 1 {
		return errors.New("the amount must be greater than 0")
	}

	if t.Amount > 1000 {
		return errors.New("limit exceeded")
	}

	return nil
}
