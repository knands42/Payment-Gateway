package process_transaction

type TransactionDTOInput struct {
	ID                        string  `json:"id"`
	AccountId                 string  `json:"account_id"`
	CreditCardNumber          string  `json:"credit_card_number"`
	CreditCardName            string  `json:"credit_card_name"`
	CreditCardExpirationMonth string  `json:"credit_card_expiration_month"`
	CreditCardExpirationYear  string  `json:"credit_card_expiration_year"`
	CreditCardCVV             string  `json:"credit_card_cvv"`
	Amount                    float64 `json:"amount"`
}

type TransactionDTOOutput struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
