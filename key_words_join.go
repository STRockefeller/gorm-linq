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

// joinSql is a function type that represents a join condition for a gorm.DB object.
// It takes a pointer to a gorm.DB object as input and returns a pointer to a gorm.DB object as output.
type joinSql func(*gorm.DB) *gorm.DB

// NewJoinOn returns a closure function that adds a join clause to a SQL query.
// The closure function takes a pointer to a 'gorm.DB' object as input and returns a pointer to the same object.
// The join clause is added using the 'Joins' method of the 'gorm.DB' object.
//
// Parameters:
//  - joinType: The type of join to be performed (e.g., InnerJoin, LeftJoin, RightJoin).
//  - joinTable: The table to be joined with.
//  - on: The condition for the join clause.
func NewJoinOn(joinType JoinType, joinTable schema.Tabler, on string) joinSql {
	return func(d *gorm.DB) *gorm.DB {
		return d.Joins(fmt.Sprintf("%s %s ON %s", joinType, joinTable.TableName(), on))
	}
}

// NewJoinQuery returns a closure function that can be used to join tables in a database query.
//
// Parameters:
//  - qs: A struct of type 'QueryString' that contains the query string and its arguments.
func NewJoinQuery(qs QueryString) joinSql {
	return func(d *gorm.DB) *gorm.DB {
		return d.Joins(qs.Query, qs.Args...)
	}
}

// Join executes a join condition on the database and returns the result and any error that occurred.
//
// Parameters:
//  - ctx: The context for the database operation.
//  - selection: A struct that contains the query string and its arguments.
//  - joinCondition: A closure function that represents the join condition for the query.
func (container DB[T]) Join(ctx context.Context, selection QueryString, joinCondition joinSql) ([]map[string]any, error) {
	var res []map[string]any
	err := joinCondition(container.db.WithContext(ctx).Select(selection.Query, selection.Args...)).
		Scan(&res).Error

	return res, err
}
