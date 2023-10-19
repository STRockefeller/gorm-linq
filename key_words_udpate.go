package linq

import "context"

// Updates updates the data in the database with the provided instance.
// It takes a context and an instance of type T as inputs and returns the number of rows affected and an error.
func (container DB[T]) Updates(ctx context.Context, instance T) (rowsAffected int64, err error) {
	res := container.db.WithContext(ctx).Updates(&instance)
	return res.RowsAffected, res.Error
}

// UpdateWithMap updates the database with the provided column values.
// It takes a context and a map of column values as inputs, and it returns the number of rows affected and an error.
//
// Parameters:
//  - ctx (context.Context): The context in which the update operation is performed.
//  - columnValues (map[string]any): A map of column names and their corresponding values to be updated in the database.
func (container DB[T]) UpdateWithMap(ctx context.Context, columnValues map[string]any) (rowsAffected int64, err error) {
	res := container.db.WithContext(ctx).Updates(columnValues)
	return res.RowsAffected, res.Error
}

// UpdateByRequest updates the database with the provided request.
// It takes a context and an UpdateRequest object as inputs and returns the number of rows affected and an error.
//
// Parameters:
//  - ctx (context.Context): The context in which the update operation is performed.
//  - req (UpdateRequest): An object that contains the necessary information for the update operation.
func (container DB[T]) UpdateByRequest(ctx context.Context, req UpdateRequest) (rowsAffected int64, err error) {
	return container.UpdateWithMap(ctx, req.Update())
}
