package linq

import "context"

func (container DB[T]) Count(ctx context.Context) (result int64, err error) {
	err = container.db.WithContext(ctx).Count(&result).Error
	return
}
