package linq

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// Updates method updates the data in the database with the provided instance.
func TestUpdatesMethodUpdatesDataWithInstance(t *testing.T) {
	assert := assert.New(t)
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	mock.ExpectBegin()
	// Set expectations on the mock DB for your test case
	mock.ExpectExec("UPDATE `test_tables` SET `id`=\\?,`column1`=\\?,`column2`=\\? WHERE `id` = \\?").WithArgs(1, "value1", "value2", 1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	rowsAffected, err := db.WhereRaw(QueryString{Query: "`id` = ?", Args: []any{1}}).Updates(ctx, TestTable{ID: 1, Column1: "value1", Column2: "value2"})

	// Assert
	assert.NoError(err)
	assert.Equal(int64(1), rowsAffected)
	assert.NoError(mock.ExpectationsWereMet())
}

// UpdateWithMap method updates the database with the provided column values.
func TestUpdateWithMapUpdatesDatabaseWithColumnValues(t *testing.T) {
	assert := assert.New(t)
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	mock.ExpectBegin()
	// Set expectations on the mock DB for your test case
	mock.ExpectExec("UPDATE `test_tables` SET `column1`=\\?,`column2`=\\? WHERE 1=1").WithArgs("value1", "value2").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	rowsAffected, err := db.WhereRaw(QueryString{Query: "1=1"}).UpdateWithMap(ctx, map[string]interface{}{"column1": "value1", "column2": "value2"})

	// Assert
	assert.NoError(err)
	assert.Equal(int64(1), rowsAffected)
	assert.NoError(mock.ExpectationsWereMet())
}

// UpdateByRequest method updates the database with the provided request.
func TestUpdateByRequestUpdatesDatabaseWithRequest(t *testing.T) {
	assert := assert.New(t)
	db, mock, _ := MockDB[TestTable]()
	ctx := context.Background()

	mock.ExpectBegin()
	// Set expectations on the mock DB for your test case
	mock.ExpectExec("UPDATE `test_tables` SET `column1`=\\?,`column2`=\\? WHERE id = 1").WithArgs("value1", "value2").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	ur := updateReq{
		query:         []QueryString{},
		updatedValues: map[string]any{},
	}
	req := ur.whereTest(QueryString{Query: "id = 1"}).updateTest("column1", "value1").updateTest("column2", "value2")
	rowsAffected, err := db.WhereByRequest(req).UpdateByRequest(ctx, req)

	// Assert
	assert.NoError(err)
	assert.Equal(int64(1), rowsAffected)
	assert.NoError(mock.ExpectationsWereMet())
}

// Updates method returns an error when the context is canceled.
func TestUpdatesMethodReturnsErrorWhenContextCanceled(t *testing.T) {
	assert := assert.New(t)
	db, _, _ := MockDB[TestTable]()
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the context

	// Act
	rowsAffected, err := db.Updates(ctx, TestTable{ID: 1, Column1: "value1", Column2: "value2"})

	// Assert
	assert.Error(err)
	assert.Equal(int64(0), rowsAffected)
}

// UpdateWithMap method returns an error when the context is canceled.
func TestUpdateWithMapReturnsErrorWhenContextCanceled(t *testing.T) {
	assert := assert.New(t)
	db, _, _ := MockDB[TestTable]()
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the context

	// Act
	rowsAffected, err := db.UpdateWithMap(ctx, map[string]interface{}{"column1": "value1", "column2": "value2"})

	// Assert
	assert.Error(err)
	assert.Equal(int64(0), rowsAffected)
}

// UpdateByRequest method returns an error when the context is canceled.
func TestUpdateByRequestReturnsErrorWhenContextCanceled(t *testing.T) {
	assert := assert.New(t)
	db, _, _ := MockDB[TestTable]()
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the context

	// Act
	ur := updateReq{
		query:         []QueryString{},
		updatedValues: map[string]any{},
	}
	rowsAffected, err := db.UpdateByRequest(ctx, ur.whereTest(QueryString{Query: "id = 1"}).updateTest("column1", "value1").updateTest("column2", "value2"))

	// Assert
	assert.Error(err)
	assert.Equal(int64(0), rowsAffected)
}
