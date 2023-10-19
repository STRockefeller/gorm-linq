package linq

import "context"

func (container DB[T]) Updates(ctx context.Context, instance T) (rowsAffected int64, err error) {
	res := container.db.WithContext(ctx).Updates(&instance)
	return res.RowsAffected, res.Error
}
