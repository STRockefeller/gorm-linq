package linq

import (
	"context"
)

// Create creates new records in the database.
func (container DB[T]) Create(ctx context.Context, instances ...T) error {
	return container.db.WithContext(ctx).Create(&instances).Error
}
