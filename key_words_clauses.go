package linq

import "gorm.io/gorm/clause"

// Clauses updates the container's db field by adding the provided conditions as clauses.
// It returns the updated container.
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

func parseLockOption(opt LockOption) []forUpdateOption {
	switch opt {
	case OptionNoWait:
		return []forUpdateOption{NoWait()}
	default:
		return []forUpdateOption{}
	}
}

// LockByRequest locks the container based on the provided LockableRequest.
// If a lock is required, it calls the ForUpdate method with the parsed lock option.
// If a lock is not required, it returns the original container.
func (container DB[T]) LockByRequest(req LockableRequest) DB[T] {
	if req.Lock() {
		return container.ForUpdate(parseLockOption(req.LockOption())...)
	}
	return container
}
