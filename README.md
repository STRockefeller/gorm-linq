# Gormlinq

Gormlinq is a library that provides a LINQ-like interface for the GORM ORM library. It allows you to write more expressive queries and perform common database operations in a more concise and readable way.

## Installation

To install Gormlinq, use `go get`:

```
go get github.com/STRockefeller/gormlinq
```

## Usage

First, import the `gormlinq` package:

```go
import "github.com/STRockefeller/gormlinq"
```

Then, create a new `gorm.DB` instance and pass it to `gormlinq.NewDB` along with the struct type you want to work with:

```go
db, err := gorm.Open(mysql.Open("dsn"), &gorm.Config{})
if err != nil {
    // handle error
}

type User struct {
    ID   int
    Name string
}

users := gormlinq.NewDB[User](db)
```

You can then use `gormlinq` methods to query the database:

```go
// Find all users where Name starts with "A"
var result linq.Linq[User]
err := users.Where(User{Name: "A%"}).Find(context.Background(), &result)

// Find the first user where ID is 42
var user User
err := users.Where(User{ID: 42}).Take(context.Background(), &user)

// Update all users where Name starts with "A"
rowsAffected, err := users.Where(User{Name: "A%"}).Updates(context.Background(), User{Name: "NewName"})

// Delete all users where Name starts with "A"
rowsAffected, err := users.Delete(context.Background(), User{Name: "A%"})

// Find all users for update
var result linq.Linq[User]
err := users.FindForUpdate(context.Background(), gormlinq.NoWait(), &result)
```

You can also chain multiple `gormlinq` methods together to build more complex queries:

```go
// Find all users where Name starts with "A" and ID is less than 100, ordered by Name
var result linq.Linq[User]
err := users.Where(User{Name: "A%"}).WhereRaw("ID < ?", 100).Order("Name").Find(context.Background(), &result)
```

`gormlinq` also supports upserts:

```go
// Upsert a single user
err := users.Upsert(context.Background(), []User{{ID: 42, Name: "NewName"}}, clause.OnConflict{
    Columns:   []clause.Column{{Name: "id"}},
    DoUpdates: clause.AssignmentColumns([]string{"name"}),
})
```

You can also use `gormlinq` with GORM's `Scope` method to apply a function to the underlying `gorm.DB` instance:

```go
users.Scope(func(db *gorm.DB) *gorm.DB {
    return db.Where("deleted_at IS NULL")
}).Find(context.Background(), &result)
```

## License

Gormlinq is released under the MIT License. See LICENSE file for details.