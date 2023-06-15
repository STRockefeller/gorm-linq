package linq

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/STRockefeller/go-linq"
)

type DB[T any] struct {
	db *gorm.DB
}

func NewDB[T any](db *gorm.DB) DB[T] {
	return DB[T]{db: db.Model(new(T))}
}

func (container DB[T]) Create(ctx context.Context, instances ...T) error {
	return container.db.WithContext(ctx).Create(&instances).Error
}

func (container DB[T]) DeleteWithCondition(ctx context.Context, condition T) (rawsAffected int64, err error) {
	res := container.db.WithContext(ctx).Delete(&condition)
	return res.RowsAffected, res.Error
}

func (container DB[T]) Delete(ctx context.Context) (rawsAffected int64, err error) {
	return container.DeleteWithCondition(ctx, *new(T))
}

func (container DB[T]) SelectRaw(selectedColumns string) DB[T] {
	container.db = container.db.Select(selectedColumns)
	return container
}

func (container DB[T]) Where(condition T) DB[T] {
	container.db = container.db.Where(&condition)
	return container
}

func (container DB[T]) WhereRaw(condition string, args ...interface{}) DB[T] {
	container.db = container.db.Where(condition, args...)
	return container
}

func (container DB[T]) Find(ctx context.Context) (result linq.Linq[T], err error) {
	err = container.db.WithContext(ctx).Find(&result).Error
	return
}

// panics if something went wrong
func (container DB[T]) FindWithoutError(ctx context.Context) (result linq.Linq[T]) {
	err := container.db.WithContext(ctx).Find(&result).Error
	if err != nil {
		panic(err.Error())
	}
	return
}

func (container DB[T]) Take(ctx context.Context) (result T, err error) {
	err = container.db.WithContext(ctx).Take(&result).Error
	return
}

func (container DB[T]) Count(ctx context.Context) (result int64, err error) {
	err = container.db.WithContext(ctx).Count(&result).Error
	return
}

func (container DB[T]) Updates(ctx context.Context, instance T) (rowsAffected int64, err error) {
	res := container.db.WithContext(ctx).Updates(&instance)
	return res.RowsAffected, res.Error
}

func (container DB[T]) ForUpdate(opts ...forUpdateOption) DB[T] {
	var opt string
	if parseForUpdateOptions(opts...).NoWait {
		opt = "NOWAIT"
	}
	container.db = container.db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Options:  opt,
	})

	return container
}

type forUpdateOptions struct {
	NoWait bool
}

type forUpdateOption func(*forUpdateOptions)

func NoWait() forUpdateOption {
	return func(fuo *forUpdateOptions) {
		fuo.NoWait = true
	}
}

func parseForUpdateOptions(opts ...forUpdateOption) (res forUpdateOptions) {
	for _, opt := range opts {
		opt(&res)
	}
	return
}

func (container DB[T]) FindForUpdate(ctx context.Context, opts ...forUpdateOption) (result linq.Linq[T], err error) {
	return container.ForUpdate(opts...).Find(ctx)
}

func (container DB[T]) TakeForUpdate(ctx context.Context, opts ...forUpdateOption) (result T, err error) {
	return container.ForUpdate(opts...).Take(ctx)
}

func (container DB[T]) Upsert(ctx context.Context, clause clause.OnConflict, instances ...T) error {
	return container.db.WithContext(ctx).Clauses(clause).Create(&instances).Error
}

func (container DB[T]) Scope(f func(*gorm.DB) *gorm.DB) DB[T] {
	container.db = f(container.db)
	return container
}

func (container DB[T]) Error() error {
	return container.db.Error
}
