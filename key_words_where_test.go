package linq

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// WhereRaw

func TestWhereRawSelectObjects(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	selectedColumns := []string{"column1", "column2"}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT `column1`,`column2` FROM `test_tables` WHERE column1 = \\?").WithArgs("hello").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.SelectRaw(selectedColumns).WhereRaw(QueryString{Query: "column1 = ?", Args: []any{"hello"}}).Find(context.Background())
	assert.NoError(err)

	// Assert
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Deleting all rows returns the number of affected rows and no error.
func TestWhereDeleteRows(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()

	mock.ExpectBegin()
	// Set expectations on the mock DB for your test case
	mock.ExpectExec("DELETE FROM `test_tables` WHERE `test_tables`.`id` = \\?").WithArgs(2).WillReturnResult(sqlmock.NewResult(1, 5))
	mock.ExpectCommit()

	// Act
	rowsAffected, err := db.Where(TestTable{ID: 2}).Delete(context.Background())
	assert.NoError(err)

	// Assert
	assert.Equal(int64(5), rowsAffected)
	assert.NoError(mock.ExpectationsWereMet())
}
