package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db}
}

func (r *TransactionRepositoryDb) Insert(id, account, status, errorMessage string, amount float64) error {
	stmt, err := r.db.Prepare("INSERT INTO transactions(id, account_id, amount, status, error_message, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, account, amount, status, errorMessage, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}
