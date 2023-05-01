package factory

import (
	"database/sql"
	tracer_adapter "github.com/caiofernandes00/payment-gateway/adapter/trace"

	repo "github.com/caiofernandes00/payment-gateway/adapter/repository"
	"github.com/caiofernandes00/payment-gateway/domain/repository"
)

type RepositoryDatabaseFactory struct {
	DB   *sql.DB
	otel tracer_adapter.TraceClosure
}

func NewRepositoryDatabaseFactory(db *sql.DB, otel tracer_adapter.TraceClosure) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{
		DB:   db,
		otel: otel,
	}
}

func (f *RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	return repo.NewTransactionRepositoryDb(f.DB, f.otel)
}
