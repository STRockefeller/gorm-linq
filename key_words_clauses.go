package linq

import "gorm.io/gorm/clause"

func (container DB[T]) Clauses(conditions ...clause.Expression) DB[T] {
	container.db = container.db.Clauses(conditions...)
	return container
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
