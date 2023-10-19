package linq

import (
	"gorm.io/gorm"
)

// DB is a generic struct that wraps around the 'gorm.DB' struct from the 'gorm' package.
type DB[T any] struct {
	db *gorm.DB
}

// NewDB creates a new instance of 'DB' by initializing the 'db' field with the provided 'gorm.DB' instance.
func NewDB[T any](db *gorm.DB) DB[T] {
	return DB[T]{db: db.Model(new(T))}
}

// Scope applies a scope to the query by calling the provided function 'f' with the 'db' field as an argument.
// It returns the modified 'DB' instance.
func (container DB[T]) Scope(f func(*gorm.DB) *gorm.DB) DB[T] {
	container.db = f(container.db)
	return container
}

// Error returns the error associated with the 'DB' instance.
func (container DB[T]) Error() error {
	return container.db.Error
}
