package repository

import (
	"context"
	"database/sql"
	tracer_adapter "github.com/caiofernandes00/payment-gateway/adapter/trace"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	migrationDir                             = os.DirFS("migrations")
	otel         tracer_adapter.TraceClosure = func(ctx context.Context, tracingName string, fn func(context.Context)) context.Context {
		fn(ctx)
		return ctx
	}
	ctx = context.Background()
)

func Test_DbInsert(t *testing.T) {
	// Arrange
	db, err := sql.Open("sqlite3", ":memory:")
	Up(db, migrationDir)
	defer Down(db, migrationDir)

	repo := NewTransactionRepositoryDb(db, otel)

	// Act
	err = repo.Insert(ctx, "id", "account", "status", "error message", 100.00)

	// Assert
	assert.Nil(t, err)
}
