package linq

import "context"

// Delete deletes all rows from the database.
// It takes a context as input and returns the number of rows affected and an error.
// This method internally calls DeleteWithCondition with a new instance of the generic type T as the condition.
func (container DB[T]) Delete(ctx context.Context) (rawsAffected int64, err error) {
	res := container.db.WithContext(ctx).Delete(*new(T))
	return res.RowsAffected, res.Error
}
