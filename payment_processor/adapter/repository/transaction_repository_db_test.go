package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var migrationDir = os.DirFS("migration")

func Test_DbInsert(t *testing.T) {
	// Arrange
	db := Up(migrationDir)
	defer Down(db, migrationDir)
	repo := NewTransactionRepositoryDb(db)

	// Act
	err := repo.Insert("id", "account", "status", "error message", 100.00)

	// Assert
	assert.Nil(t, err)
}
