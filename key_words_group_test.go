package linq

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/STRockefeller/go-linq"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Group elements in database container based on specified columns and aggregation functions
func TestGroupElementsBasedOnColumns(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	type selectedValues struct {
		Column1 string
		Count   int
	}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT `column1`,count\\(column2\\) FROM `test_tables` GROUP BY `column1` HAVING count > 5").WillReturnRows(sqlmock.NewRows([]string{"column1", "count"}).AddRow("hello", 8))

	// Act
	var result selectedValues
	db.SelectRaw([]string{"column1", "count(column2)"}).GroupByHaving([]string{"column1"}, QueryString{Query: "count > 5"}).Scope(func(d *gorm.DB) *gorm.DB {
		return d.Find(&result)
	})

	// Assert
	assert.Equal(selectedValues{Column1: "hello", Count: 8}, result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Group elements in database container with one element
func TestGroupElementsWithOneElement(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	columns := []string{"column1", "column2"}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables` GROUP BY column1, column2 LIMIT 1").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.GroupBy(columns).Take(context.Background())

	// Assert
	assert.NoError(err)
	assert.Equal(TestTable{Column1: "hello", Column2: "world"}, result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Group elements in database container with multiple elements but only one distinct value for specified columns
func TestGroupElementsWithOneDistinctValue(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	columns := []string{"column1", "column2"}

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT \\* FROM `test_tables` GROUP BY column1, column2").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world").AddRow("hello", "kitty"))

	// Act
	result, err := db.GroupBy(columns).Find(context.Background())

	// Assert
	assert.NoError(err)
	assert.Equal(linq.NewLinq([]TestTable{{Column1: "hello", Column2: "world"}, {Column1: "hello", Column2: "kitty"}}), result)
	assert.NoError(mock.ExpectationsWereMet())
}
