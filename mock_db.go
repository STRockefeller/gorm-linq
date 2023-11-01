package linq

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MockDB[T any]() (*DB[T], sqlmock.Sqlmock, error) {
	// Create a mock DB connection using go-sqlmock
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	rows := sqlmock.NewRows([]string{"VERSION()"}).AddRow("1.2.3")
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(rows)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	// Create a new instance of DB and initialize the 'db' field with the gorm.DB instance
	db := NewDB[T](gormDB)

	return &db, mock, nil
}
