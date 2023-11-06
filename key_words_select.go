package linq

import (
	"context"

	"github.com/STRockefeller/go-linq"
)

// SelectRaw is a method defined in the DB struct. It takes a string slice parameter selectedColumns and returns the modified DB object.
func (container DB[T]) SelectRaw(selectedColumns []string) DB[T] {
	container.db = container.db.Select(selectedColumns)
	return container
}

func (container DB[T]) Find(ctx context.Context) (result linq.Linq[T], err error) {
	items := []T{}
	err = container.db.WithContext(ctx).Find(&items).Error
	result = linq.New(items)
	return
}

// panics if something went wrong
func (container DB[T]) FindWithoutError(ctx context.Context) (result linq.Linq[T]) {
	items := []T{}
	err := container.db.WithContext(ctx).Find(&items).Error
	if err != nil {
		panic(err.Error())
	}
	result = linq.New(items)
	return
}

func (container DB[T]) Take(ctx context.Context) (result T, err error) {
	err = container.db.WithContext(ctx).Take(&result).Error
	return
}

func (container DB[T]) FindForUpdate(ctx context.Context, opts ...forUpdateOption) (result linq.Linq[T], err error) {
	return container.ForUpdate(opts...).Find(ctx)
}

func (container DB[T]) TakeForUpdate(ctx context.Context, opts ...forUpdateOption) (result T, err error) {
	return container.ForUpdate(opts...).Take(ctx)
}

func (container DB[T]) Distinct(columns []string) DB[T] {
	container.db = container.db.Distinct(linq.Select(columns, func(c string) any { return c }).ToSlice()...)
	return container
}
