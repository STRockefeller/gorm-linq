package linq

import "context"

// Count is a method that belongs to the DB type in the linq package.
// It takes a context.Context as input and returns the count of records in the database and any error that occurred during the count operation.
func (container DB[T]) Count(ctx context.Context) (number int64, err error) {
	err = container.db.WithContext(ctx).Count(&number).Error
	return
}
