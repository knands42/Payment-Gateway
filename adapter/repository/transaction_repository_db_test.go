package repository

import (
	"os"
	"testing"

	"github.com/caiofernandes00/payment-gateway/adapter/repository/fixture"
	"github.com/stretchr/testify/assert"
)

var migrationDir = os.DirFS("fixture/migration")

func Test_DbInsert(t *testing.T) {
	// Arrange
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)
	repo := NewTransactionRepositoryDb(db)

	// Act
	err := repo.Insert("id", "account", "status", "error message", 100.00)

	// Assert
	assert.Nil(t, err)
}
