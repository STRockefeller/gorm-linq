# Gorm-linq

![license](https://img.shields.io/github/license/STRockefeller/gorm-linq?style=plastic)![code size](https://img.shields.io/github/languages/code-size/STRockefeller/gorm-linq?style=plastic)![open issues](https://img.shields.io/github/issues/STRockefeller/gorm-linq?style=plastic)![closed issues](https://img.shields.io/github/issues-closed/STRockefeller/gorm-linq?style=plastic)![go version](https://img.shields.io/github/go-mod/go-version/STRockefeller/gorm-linq?style=plastic)![latest version](https://img.shields.io/github/v/tag/STRockefeller/gorm-linq?style=plastic)[![Go Report Card](https://goreportcard.com/badge/github.com/STRockefeller/gorm-linq)](https://goreportcard.com/report/github.com/STRockefeller/gorm-linq)[![Coverage Status](https://coveralls.io/repos/github/STRockefeller/gorm-linq/badge.svg)](https://coveralls.io/github/STRockefeller/gorm-linq)

Gorm-linq is a library that provides a LINQ-like interface for the GORM ORM library. It allows you to write more expressive queries and perform common database operations in a more concise and readable way.

## Installation

To install Gorm-linq, use `go get`:

```
go get github.com/STRockefeller/gorm-linq
```

## Usage

First, import the `linq` package:

```go
import "github.com/STRockefeller/gorm-linq"
```

Then, create a new `gorm.DB` instance and pass it to `linq.NewDB` along with the struct type you want to work with:

```go
db, err := gorm.Open(mysql.Open("dsn"), &gorm.Config{})
if err != nil {
    // handle error
}

type User struct {
    ID   int
    Name string
}

users := linq.NewDB[User](db)
```

You can then use `gorm-linq` methods to query the database:

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
err := users.FindForUpdate(context.Background(), linq.NoWait(), &result)
```

You can also chain multiple `gorm-linq` methods together to build more complex queries:

```go
// Find all users where Name starts with "A" and ID is less than 100, ordered by Name
var result linq.Linq[User]
err := users.Where(User{Name: "A%"}).WhereRaw("ID < ?", 100).Order("Name").Find(context.Background(), &result)
```

`gorm-linq` also supports upserts:

```go
// Upsert a single user
err := users.Upsert(context.Background(), []User{{ID: 42, Name: "NewName"}}, clause.OnConflict{
    Columns:   []clause.Column{{Name: "id"}},
    DoUpdates: clause.AssignmentColumns([]string{"name"}),
})
```

You can also use `gorm-linq` with GORM's `Scope` method to apply a function to the underlying `gorm.DB` instance:

```go
users.Scope(func(db *gorm.DB) *gorm.DB {
    return db.Where("deleted_at IS NULL")
}).Find(context.Background(), &result)
```

## License

Gorm-linq is released under the MIT License. See LICENSE file for details.
