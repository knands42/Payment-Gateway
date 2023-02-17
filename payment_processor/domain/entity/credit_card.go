package entity

import (
	"errors"
	"regexp"

	NotificationPackage "github.com/caiofernandes00/payment-gateway/domain/notification"
)

const (
	CREDIT_CONTEXT = "creditcard"

	STATUS_APPROVED = "approved"
	STATUS_REJECTED = "rejected"
)

type CreditCard struct {
	Number          string
	Name            string
	ExpirationMonth int
	ExpirationYear  int
	CVV             string
	notification    NotificationPackage.Notification
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
	cc.validateNumber()
	cc.validateExpirationMonth()
	cc.validateExpirationYear()
	cc.validateCVV()

	if cc.notification.HasErrors() {
		return errors.New(cc.notification.Messages(CREDIT_CONTEXT))
	}

	return nil
}

func (cc *CreditCard) validateNumber() {
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
		cc.notification.AddError("invalid credit card number", CREDIT_CONTEXT)
	}
}

func (cc *CreditCard) validateExpirationMonth() {
	isValid := cc.ExpirationMonth >= 1 && cc.ExpirationMonth <= 12
	if !isValid {
		cc.notification.AddError("invalid credit card expiration month", CREDIT_CONTEXT)
	}
}

func (cc *CreditCard) validateExpirationYear() {
	isValid := cc.ExpirationYear >= 2020
	if !isValid {
		cc.notification.AddError("invalid credit card expiration year", CREDIT_CONTEXT)
	}
}

func (cc *CreditCard) validateCVV() {
	isValid := len(cc.CVV) == 3 || len(cc.CVV) == 4
	if !isValid {
		cc.notification.AddError("invalid credit card cvv", CREDIT_CONTEXT)
	}
}
