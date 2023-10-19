package linq

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type JoinType string

const (
	Unspecified JoinType = ""
	LeftJoin    JoinType = "LEFT JOIN"
	RightJoin   JoinType = "RIGHT JOIN"
	InnerJoin   JoinType = "INNER JOIN"
)

type joinSql func(*gorm.DB) *gorm.DB

func NewJoinOn(joinType JoinType, joinTable schema.Tabler, on string) joinSql {
	return func(d *gorm.DB) *gorm.DB {
		return d.Joins(fmt.Sprintf("%s %s ON %s", joinType, joinTable.TableName(), on))
	}
}

func NewJoinQuery(qs QueryString) joinSql {
	return func(d *gorm.DB) *gorm.DB {
		return d.Joins(qs.Query, qs.Args...)
	}
}

func (container DB[T]) Join(ctx context.Context, selection QueryString, joinCondition joinSql) ([]map[string]any, error) {
	var res []map[string]any
	err := joinCondition(container.db.WithContext(ctx).Select(selection.Query, selection.Args...)).
		Scan(&res).Error

	return res, err
}
