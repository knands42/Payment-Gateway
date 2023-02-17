package factory

import "github.com/caiofernandes00/payment-gateway/domain/repository"

type TransactionRepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
