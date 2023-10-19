package linq

import "fmt"

type QueryString struct {
	Query string
	Args  []any
}

func WhereInCondition[T any](columnName string) func([]T) QueryString {
	return func(items []T) QueryString {
		if columnName == "" {
			panic("columnName cannot be empty")
		}
		return QueryString{
			Query: fmt.Sprintf("%s IN ?", columnName),
			Args:  []any{items},
		}
	}
}

func NewQueryString(sql string, vars ...any) QueryString {
	return QueryString{
		Query: sql,
		Args:  vars,
	}
}
