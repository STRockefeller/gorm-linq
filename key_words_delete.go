package linq

import "context"

func (container DB[T]) DeleteWithCondition(ctx context.Context, condition T) (rawsAffected int64, err error) {
	res := container.db.WithContext(ctx).Delete(&condition)
	return res.RowsAffected, res.Error
}

func (container DB[T]) Delete(ctx context.Context) (rawsAffected int64, err error) {
	return container.DeleteWithCondition(ctx, *new(T))
}
