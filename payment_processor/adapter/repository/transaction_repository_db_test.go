package repository

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var migrationDir = os.DirFS("migrations")

func Test_DbInsert(t *testing.T) {
	// Arrange
	db, err := sql.Open("sqlite3", ":memory:")
	Up(db, migrationDir)
	defer Down(db, migrationDir)
	repo := NewTransactionRepositoryDb(db)

	// Act
	err = repo.Insert("id", "account", "status", "error message", 100.00)

	// Assert
	assert.Nil(t, err)
}
