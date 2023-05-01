package repository

import "context"

type TransactionRepository interface {
	Insert(ctx context.Context, id, account, status, errorMessage string, amount float64) error
}
