package linq

import "context"

// DeleteWithCondition deletes rows from the database based on a given condition.
// It takes a context and a condition as inputs and returns the number of rows affected and an error.
func (container DB[T]) DeleteWithCondition(ctx context.Context, condition T) (rawsAffected int64, err error) {
	res := container.db.WithContext(ctx).Delete(&condition)
	return res.RowsAffected, res.Error
}

// Delete deletes all rows from the database.
// It takes a context as input and returns the number of rows affected and an error.
// This method internally calls DeleteWithCondition with a new instance of the generic type T as the condition.
func (container DB[T]) Delete(ctx context.Context) (rawsAffected int64, err error) {
	return container.DeleteWithCondition(ctx, *new(T))
}
