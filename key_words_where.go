package linq

// Where applies a condition to the db field of the container object and returns the modified container.
func (container DB[T]) Where(condition T) DB[T] {
	container.db = container.db.Where(&condition)
	return container
}

// WhereRaw applies a query string to the db field of the container object.
func (container DB[T]) WhereRaw(qs QueryString) DB[T] {
	container.db = container.db.Where(qs.Query, qs.Args...)
	return container
}

// WhereByRequest applies a series of WhereRaw queries to the container's db field based on the conditions specified in the QueryRequest.
//
// Parameters:
//  - req: A QueryRequest object that contains a function called Where which returns an array of QueryString objects. Each QueryString object represents a condition to be applied to the container's db field.
func (container DB[T]) WhereByRequest(req QueryRequest) DB[T] {
	for _, qs := range req.Where() {
		container.WhereRaw(qs)
	}
	return container
}
