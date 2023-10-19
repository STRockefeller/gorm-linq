package linq

import (
	"strings"
)

func (container DB[T]) GroupBy(columns []string) DB[T] {
	container.db = container.db.Group(strings.Join(columns, ", "))
	return container
}

func (container DB[T]) GroupByHaving(columns []string, condition QueryString) DB[T] {
	container.db = container.db.Group(strings.Join(columns, ", ")).Having(condition.Query, condition.Args...)
	return container
}
