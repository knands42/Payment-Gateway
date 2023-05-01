package repository

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	migrationDir = os.DirFS("migrations")
	ctx          = context.Background()
	otel         = func(ctx context.Context, name string, f func()) {
		f()
	}
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
