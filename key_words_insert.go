package linq

import (
	"context"

	"gorm.io/gorm/clause"
)

// Create creates new records in the database.
func (container DB[T]) Create(ctx context.Context, instances ...T) error {
	return container.db.WithContext(ctx).Create(&instances).Error
}

// Upsert performs an upsert operation in the database.
func (container DB[T]) Upsert(ctx context.Context, clause clause.OnConflict, instances ...T) error {
	return container.db.WithContext(ctx).Clauses(clause).Create(&instances).Error
}
