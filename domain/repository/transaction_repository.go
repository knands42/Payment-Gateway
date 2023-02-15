package repository

type TransactionRepository interface {
	Insert(id, account, status, errorMessage string, amount float64) error
}
