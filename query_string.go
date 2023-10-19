package linq

import "fmt"

// QueryString is a struct that represents a SQL query string and its arguments.
type QueryString struct {
	Query string
	Args  []any
}

func NewQueryString(sql string, vars ...any) QueryString {
	return QueryString{
		Query: sql,
		Args:  vars,
	}
}

// The `WhereInCondition` function returns a closure that takes a slice of any type `T` as input and returns a `QueryString` struct.
// The closure constructs a SQL query string with the provided `columnName` and the input items as arguments.
// If the `columnName` is empty, the function panics with an error message.
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
