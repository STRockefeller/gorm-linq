package linq

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// Successfully creates new records in the database with valid input
func TestCreatesNewRecordsWithValidInput(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	instances := []TestTable{{
		ID:      1,
		Column1: "value1",
		Column2: "value2",
	}}

	mock.ExpectBegin()
	// Set expectations on the mock DB for your test case
	mock.ExpectExec("INSERT INTO `test_tables` \\(`column1`,`column2`,`id`\\) VALUES \\(\\?,\\?,\\?\\)").WithArgs("value1", "value2", 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	err := db.Create(context.Background(), instances...)

	// Assert
	assert.NoError(err)
	assert.NoError(mock.ExpectationsWereMet())
}

// Returns error when input is empty
func TestReturnsErrorWhenInputIsEmpty(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	mock.ExpectBegin()
	mock.ExpectRollback()
	// Act
	err := db.Create(context.Background())

	// Assert
	assert.EqualError(err, "empty slice found")
	assert.NoError(mock.ExpectationsWereMet())
}
