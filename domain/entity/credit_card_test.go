package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreditCardNumber(t *testing.T) {
	// Invalid single error
	_, err := NewCreditCard("1234567890123456", "John Doe", 1, 2020, "123")
	assert.Equal(t, "creditcard: invalid credit card number", err.Error())

	// Create a Visa card
	visa, err := NewCreditCard("4111111111111111", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, "4111111111111111", visa.Number)

	// Create a Master card
	master, err := NewCreditCard("5555555555554444", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, "5555555555554444", master.Number)

	// Create a Amex card
	amex, err := NewCreditCard("378282246310005", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, "378282246310005", amex.Number)

	// Create a Discover card
	discover, err := NewCreditCard("6011111111111117", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, "6011111111111117", discover.Number)

	// Create a Diners card
	diners, err := NewCreditCard("30569309025904", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, "30569309025904", diners.Number)

	// Create a JCB card
	jcb, err := NewCreditCard("3530111333300000", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, "3530111333300000", jcb.Number)

	// Create a JCB card
	jcb2, err := NewCreditCard("3566002020360505", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, "3566002020360505", jcb2.Number)
}

func Test_CreditCardExpirationMonth(t *testing.T) {
	// Valid
	visa, err := NewCreditCard("4111111111111111", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, 1, visa.ExpirationMonth)

	// Invalid single error
	_, err = NewCreditCard("4111111111111111", "John Doe", 13, 2020, "123")
	assert.Equal(t, "creditcard: invalid credit card expiration month", err.Error())
}

func Test_CreditCardExpirationYear(t *testing.T) {
	// Valid
	visa, err := NewCreditCard("4111111111111111", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, 2020, visa.ExpirationYear)

	// Invalid single error
	_, err = NewCreditCard("4111111111111111", "John Doe", 1, 2019, "123")
	assert.Equal(t, "creditcard: invalid credit card expiration year", err.Error())
}

func Test_CreditCardCVV(t *testing.T) {
	// Valid
	visa, err := NewCreditCard("4111111111111111", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, "123", visa.CVV)

	// Invalid single error
	_, err = NewCreditCard("4111111111111111", "John Doe", 1, 2020, "1234")
	assert.Equal(t, "creditcard: invalid credit card cvv", err.Error())
}

func Test_CreditCard(t *testing.T) {
	// Valid
	cc, err := NewCreditCard("4111111111111111", "John Doe", 1, 2020, "123")
	assert.Nil(t, err)
	assert.Equal(t, "4111111111111111", cc.Number)
	assert.Equal(t, "John Doe", cc.Name)
	assert.Equal(t, 1, cc.ExpirationMonth)
	assert.Equal(t, 2020, cc.ExpirationYear)
	assert.Equal(t, "123", cc.CVV)

	// Invalid multiple errors
	_, err = NewCreditCard("4111111111111111", "John Doe", 13, 2019, "1234")
	assert.Equal(t, "creditcard: invalid credit card expiration month,invalid credit card expiration year,invalid credit card cvv", err.Error())
}
