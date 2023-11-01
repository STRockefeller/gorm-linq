package linq

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Returns a new instance of 'DB' with the provided 'gorm.DB' instance.
func TestNewDB_ReturnsNewInstanceWithProvidedGormDB(t *testing.T) {
	assert := assert.New(t)
	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"VERSION()"}).AddRow("1.2.3")
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(rows)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(err)

	result := NewDB[int](gormDB)

	assert.NotNil(result)
	assert.IsType(DB[int]{}, result)
}

// Returns an empty 'DB' instance if the provided 'gorm.DB' instance is nil.
func TestNewDB_ReturnsEmptyInstanceIfProvidedGormDBIsNil(t *testing.T) {
	assert := assert.New(t)

	result := NewDB[int](nil)

	assert.NotNil(result)
	assert.IsType(DB[int]{}, result)
	assert.Nil(result.db)
}

// Calling Scope with a valid function modifies the 'db' field and returns the modified 'DB' instance.
func TestCallingScopeWithValidFunctionModifiesDBField(t *testing.T) {
	assert := assert.New(t)
	// Arrange
	db, _, _ := MockDB[int]()
	f := func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", "John")
	}

	// Act
	result := db.Scope(f)

	// Assert
	assert.Equal(f(db.db), result.db)
}

// Calling Scope with a nil function does not modify the 'db' field and returns the original 'DB' instance.
func TestCallingScopeWithNilFunctionDoesNotModifyDBField(t *testing.T) {
	assert := assert.New(t)
	// Arrange
	db, _, _ := MockDB[int]()

	// Act
	result := db.Scope(nil)

	// Assert
	assert.Equal(db.db, result.db)
}

// Calling Error returns the error associated with the 'DB' instance.
func TestCallingErrorReturnsAssociatedError(t *testing.T) {
	assert := assert.New(t)
	// Arrange
	db, _, _ := MockDB[int]()
	err := errors.New("test error")
	db.db.Error = err

	// Act
	result := db.Error()

	// Assert
	assert.Equal(err, result)
}
