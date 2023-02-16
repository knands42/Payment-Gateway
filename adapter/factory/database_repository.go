package factory

import (
	"database/sql"

	repo "github.com/caiofernandes00/payment-gateway/adapter/repository"
	"github.com/caiofernandes00/payment-gateway/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{
		DB: db,
	}
}

func (f *RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDb(f.DB)
}
