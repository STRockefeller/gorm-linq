package linq

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// SelectRaw method modifies DB object with selected columns

func TestSelectRawModifiesDBObject(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	selectedColumns := []string{"column1", "column2"}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT `column1`,`column2` FROM `test_tables`").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.SelectRaw(selectedColumns).Find(context.Background())
	assert.NoError(err)

	// Assert
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Find method returns linq.Linq[T] and error
func TestFindMethodReturnsLinqAndError(t *testing.T) {
	assert := assert.New(t)
	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()
	expectedError := errors.New("error")
	mock.ExpectQuery("").WillReturnError(expectedError)

	// Act
	result, err := db.Find(ctx)

	// Assert
	assert.Equal(expectedError, err)
	assert.Nil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// FindWithoutError method returns linq.Linq[T] and panics if error occurs
func TestFindWithoutError(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables`").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result := db.FindWithoutError(ctx)

	// Assert
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Take method returns T and error
func TestTakeReturnsTAndError(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables` LIMIT 1").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.Take(ctx)

	// Assert
	assert.Equal("hello", result.Column1)
	assert.Equal("world", result.Column2)
	assert.NoError(err)
	assert.NoError(mock.ExpectationsWereMet())
}

// FindForUpdate method returns linq.Linq[T] and error
func TestFindForUpdateReturnsLinqAndError(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables` FOR UPDATE").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.FindForUpdate(ctx)

	// Assert
	assert.NoError(err)
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// TakeForUpdate method returns T and error
func TestTakeForUpdateReturnsTAndError(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables` LIMIT 1 FOR UPDATE").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.TakeForUpdate(ctx)
	assert.NoError(err)

	// Assert
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Distinct method modifies DB object with distinct columns
func TestDistinctModifiesDBObject(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	columns := []string{"column1", "column2"}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT DISTINCT `column1`,`column2` FROM `test_tables`").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.Distinct(columns).Find(context.Background())
	assert.NoError(err)

	// Assert
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Empty selectedColumns parameter passed to SelectRaw method
func TestSelectRawWithEmptySelectedColumns(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	selectedColumns := []string{}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.SelectRaw(selectedColumns).Find(context.Background())
	assert.NoError(err)

	// Assert
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Error occurs during DB selection in Find method
func TestErrorDuringDBSelectionInFindMethod(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables`").WillReturnError(errors.New("error occurred during DB selection"))

	// Act
	_, err := db.Find(ctx)

	// Assert
	assert.Error(err)
	assert.NoError(mock.ExpectationsWereMet())
}

// Error occurs during DB selection in FindWithoutError method
func TestFindWithoutErrorWithError(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables`").WillReturnError(errors.New("error occurred"))

	// Act and Assert
	assert.Panics(func() {
		db.FindWithoutError(ctx)
	})
	assert.NoError(mock.ExpectationsWereMet())
}

// Error occurs during DB selection in Take method
func TestErrorOccursDuringDBSelectionInTakeMethod(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables`").WillReturnError(errors.New("error occurred during DB selection"))

	// Act
	_, err := db.Take(ctx)

	// Assert
	assert.Error(err)
	assert.Equal("error occurred during DB selection", err.Error())
	assert.NoError(mock.ExpectationsWereMet())
}

// Error occurs during DB selection in FindForUpdate method
func TestFindForUpdateWithError(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()
	opts := []forUpdateOption{}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables` FOR UPDATE").WillReturnError(errors.New("error occurred"))

	// Act
	result, err := db.FindForUpdate(ctx, opts...)

	// Assert
	assert.Error(err)
	assert.Nil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Error occurs during DB selection in TakeForUpdate method
func TestErrorOccursDuringDBSelectionInTakeForUpdateMethod(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	opts := []forUpdateOption{}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables` LIMIT 1 FOR UPDATE").WillReturnError(errors.New("error occurred during DB selection"))

	// Act
	_, err := db.TakeForUpdate(context.Background(), opts...)

	// Assert
	assert.Error(err)
	assert.Equal("error occurred during DB selection", err.Error())
	assert.NoError(mock.ExpectationsWereMet())
}

// Empty columns parameter passed to Distinct method
func TestDistinctEmptyColumnsParameter(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	columns := []string{}

	// Set expectations on the mock DB for your test case

	// Act
	result := db.Distinct(columns)

	// Assert
	assert.Equal(*db, result)
	assert.NoError(mock.ExpectationsWereMet())
}
