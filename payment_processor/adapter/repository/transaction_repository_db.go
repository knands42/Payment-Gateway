package repository

import (
	"context"
	"database/sql"
	tracer_adapter "github.com/caiofernandes00/payment-gateway/adapter/trace"
	"time"
)

type TransactionRepositoryDb struct {
	db   *sql.DB
	otel tracer_adapter.TraceClosure
}

func NewTransactionRepositoryDb(db *sql.DB, otel tracer_adapter.TraceClosure) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db, otel}
}

func (r *TransactionRepositoryDb) Insert(ctx context.Context, id, account, status, errorMessage string, amount float64) error {
	var err error

	r.otel(ctx, "insert-db", func(ctx context.Context) {
		stmt, err := r.db.Prepare("INSERT INTO transactions(id, account_id, amount, status, error_message, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			return
		}

		_, err = stmt.Exec(id, account, amount, status, errorMessage, time.Now(), time.Now())
		if err != nil {
			return
		}
	})

	if err != nil {
		return err
	}

	return nil
}
