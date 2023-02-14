package entity

import (
	"errors"
	"regexp"
)

type CreditCard struct {
	Number          string
	Name            string
	ExpirationMonth int
	ExpirationYear  int
	CVV             string
}

func NewCreditCard(number, name string, expirationMonth, expirationYear int, cvv string) (*CreditCard, error) {
	cc := &CreditCard{
		Number:          number,
		Name:            name,
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
		CVV:             cvv,
	}

	if err := cc.IsValid(); err != nil {
		return nil, err
	}
	return cc, nil
}

func (cc *CreditCard) IsValid() error {
	if !cc.validateCVV() {
		return errors.New("invalid credit card cvv")
	}

	if !cc.validateExpirationMonth() {
		return errors.New("invalid credit card expiration month")
	}

	if !cc.validateExpirationYear() {
		return errors.New("invalid credit card expiration year")
	}

	if err := cc.validateNumber(); err != nil {
		return err
	}

	return nil
}

func (cc *CreditCard) validateNumber() error {
	visaRg := regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`)
	masterRg := regexp.MustCompile(`^5[1-5][0-9]{14}$`)
	amexRg := regexp.MustCompile(`^3[47][0-9]{13}$`)
	discoverRg := regexp.MustCompile(`^6(?:011|5[0-9]{2})[0-9]{12}$`)
	dinersRg := regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`)
	jcbRg := regexp.MustCompile(`^(?:2131|1800|35\d{3})\d{11}$`)

	if !visaRg.MatchString(cc.Number) &&
		!masterRg.MatchString(cc.Number) &&
		!amexRg.MatchString(cc.Number) &&
		!discoverRg.MatchString(cc.Number) &&
		!dinersRg.MatchString(cc.Number) &&
		!jcbRg.MatchString(cc.Number) {
		return errors.New("invalid credit card number")
	}

	return nil
}

func (cc *CreditCard) validateExpirationMonth() bool {
	return cc.ExpirationMonth >= 1 && cc.ExpirationMonth <= 12
}

func (cc *CreditCard) validateExpirationYear() bool {
	return cc.ExpirationYear >= 2020
}

func (cc *CreditCard) validateCVV() bool {
	return len(cc.CVV) == 3
}
