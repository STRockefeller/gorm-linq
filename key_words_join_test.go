package linq

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type joinTable struct{}

func (joinTable) TableName() string { return "join_table" }

// Adding a join clause to a SQL query using NewJoinOn function with InnerJoin, LeftJoin, and RightJoin join types
func TestNewJoinOnWithRightJoin(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	joinType := RightJoin
	joinTable := joinTable{}
	on := "test_tables.column1 = join_table.column1"

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT FROM `test_tables` RIGHT JOIN join_table ON test_tables.column1 = join_table.column1").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	joinCondition := NewJoinOn(joinType, joinTable, on)
	result, err := db.Join(context.Background(), QueryString{}, joinCondition)
	assert.NoError(err)

	// Assert
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Joining tables in a database query using NewJoinQuery function
func TestJoinTablesInDatabaseQuery(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()
	selection := QueryString{
		Query: "test_tables.column1, join_table.column1",
		Args:  []interface{}{},
	}
	joinCondition := NewJoinQuery(QueryString{Query: "test_tables.column1 = join_table.column1"})

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT test_tables.column1, join_table.column1 FROM `test_tables` test_tables.column1 = join_table.column1").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.Join(ctx, selection, joinCondition)

	// Assert
	assert.NoError(err)
	assert.NotNil(result)
	assert.NoError(mock.ExpectationsWereMet())
}

// Executing a join condition on the database using Join function
func TestJoinConditionOnDatabase(t *testing.T) {
	assert := assert.New(t)

	// Arrange
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()
	selection := QueryString{
		Query: "test_tables.column1, join_table.column1",
		Args:  []interface{}{},
	}
	joinCondition := NewJoinOn(InnerJoin, &joinTable{}, "condition")

	// Set expectations on the mock DB for your test case
	mock.ExpectQuery("SELECT test_tables.column1, join_table.column1 FROM `test_tables` INNER JOIN join_table ON condition").WillReturnRows(sqlmock.NewRows([]string{"column1", "column2"}).AddRow("hello", "world"))

	// Act
	result, err := db.Join(ctx, selection, joinCondition)

	// Assert
	assert.NotNil(result)
	assert.NoError(err)
	assert.NoError(mock.ExpectationsWereMet())
}
