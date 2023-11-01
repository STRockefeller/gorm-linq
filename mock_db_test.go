package linq

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB returns a DB instance with a non-nil db field when passed a non-nil gorm.DB instance
func Test_new_db_with_non_nil_gorm_db(t *testing.T) {
	// Create a mock gorm.DB instance
	mockDB, _, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDB,
	}), &gorm.Config{})

	// Call the code under test
	db := NewDB[int](gormDB)

	// Assert that the db field is not nil
	if db.db == nil {
		t.Errorf("Expected db field to be non-nil, but got nil")
	}
}

// MockDB returns a non-nil DB instance, sqlmock.Sqlmock instance, and nil error when called
func Test_mock_db_returns_non_nil_instances_and_nil_error(t *testing.T) {
	// Call the code under test
	db, mock, err := MockDB[int]()

	// Assert that the returned instances are not nil
	if db == nil {
		t.Errorf("Expected DB instance to be non-nil, but got nil")
	}
	if mock == nil {
		t.Errorf("Expected sqlmock.Sqlmock instance to be non-nil, but got nil")
	}

	// Assert that the error is nil
	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

// MockDB returns a DB instance with a non-nil db field when called
func Test_mock_db_returns_db_with_non_nil_db_field(t *testing.T) {
	// Call the code under test
	db, _, _ := MockDB[int]()

	// Assert that the db field is not nil
	if db.db == nil {
		t.Errorf("Expected db field to be non-nil, but got nil")
	}
}

// NewDB returns a DB instance with a nil db field when passed a nil gorm.DB instance
func Test_new_db_with_nil_gorm_db(t *testing.T) {
	// Call the code under test
	db := NewDB[int](nil)

	// Assert that the db field is nil
	if db.db != nil {
		t.Errorf("Expected db field to be nil, but got non-nil")
	}
}
