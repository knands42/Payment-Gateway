package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TransactionMultipleErrors(t *testing.T) {
	_, err := NewTransaction("", "", 0)
	assert.Equal(t, "transaction: the amount must be greater than 0", err.Error())
}

func Test_TrasactionAmount(t *testing.T) {
	// valid
	_, err := NewTransaction("123", "123", 100)
	assert.Nil(t, err)

	// invalid single error
	_, err = NewTransaction("123", "123", 0)
	assert.Equal(t, "transaction: the amount must be greater than 0", err.Error())

	_, err = NewTransaction("123", "123", 1001)
	assert.Equal(t, "transaction: limit exceeded", err.Error())
}

func Test_SetCreditCard(t *testing.T) {
	cc, _ := NewCreditCard("4111111111111111", "John Doe", 1, 2021, "123")
	transaction, _ := NewTransaction("123", "123", 100)
	transaction.SetCreditCard(*cc)
	assert.Equal(t, cc, &transaction.CreditCard)
}
