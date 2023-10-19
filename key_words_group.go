package linq

import (
	"strings"
)

// GroupBy groups the elements in the database container based on the specified columns.
// It takes a slice of strings representing the columns to group by.
// It returns the updated database container after grouping the elements based on the specified columns.
func (container DB[T]) GroupBy(columns []string) DB[T] {
	container.db = container.db.Group(strings.Join(columns, ", "))
	return container
}

// GroupByHaving groups the elements in the database container based on the specified columns and condition.
// It takes a slice of strings representing the columns to group by and a QueryString representing the condition.
// It returns the updated database container after grouping the elements based on the specified columns and condition.
func (container DB[T]) GroupByHaving(columns []string, condition QueryString) DB[T] {
	container.db = container.db.Group(strings.Join(columns, ", ")).Having(condition.Query, condition.Args...)
	return container
}
