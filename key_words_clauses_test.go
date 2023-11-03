package linq

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

// Successfully creates new records in the database with clauses
func TestUpsertsNewRecordsWithValidInput(t *testing.T) {
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
	mock.ExpectExec("INSERT INTO `test_tables` \\(`column1`,`column2`,`id`\\) VALUES \\(\\?,\\?,\\?\\) ON DUPLICATE KEY UPDATE `column1`=VALUES\\(`column1`\\),`column2`=VALUES\\(`column2`\\)").WithArgs("value1", "value2", 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	err := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(context.Background(), instances...)

	// Assert
	assert.NoError(err)
	assert.NoError(mock.ExpectationsWereMet())
}

// Select For Update method modifies DB object with selected columns

func TestSelectForUpdateModifiesDBObject(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	selectedColumns := []string{"column1", "column2"}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT `column1`,`column2` FROM `test_tables` FOR UPDATE").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.SelectRaw(selectedColumns).ForUpdate().Find(context.Background())
	assert.NoError(err)

	// Assert
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}
