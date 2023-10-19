package linq

func (container DB[T]) Where(condition T) DB[T] {
	container.db = container.db.Where(&condition)
	return container
}

func (container DB[T]) WhereRaw(qs QueryString) DB[T] {
	container.db = container.db.Where(qs.Query, qs.Args...)
	return container
}

func (container DB[T]) WhereByRequest(req QueryRequest) DB[T] {
	for _, qs := range req.Where() {
		container.WhereRaw(qs)
	}
	return container
}
