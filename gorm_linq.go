package linq

import (
	"gorm.io/gorm"
)

type DB[T any] struct {
	db *gorm.DB
}

func NewDB[T any](db *gorm.DB) DB[T] {
	return DB[T]{db: db.Model(new(T))}
}

func (container DB[T]) Scope(f func(*gorm.DB) *gorm.DB) DB[T] {
	container.db = f(container.db)
	return container
}

func (container DB[T]) Error() error {
	return container.db.Error
}
