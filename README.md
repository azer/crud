## crud

A minimalistic database library for Go, with simple and familiar interface.
It embeds [sqlx](https://github.com/jmoiron/sqlx).

Manual:

* [Install](#install)
* [Initialize](#initialize)
* [Define](#define)
* CRUD:
  * [Create](#create)
  * [Read](#read)
  * [Update](#update)
  * [Delete](#delete)
* [Tables](#tables)
* [Logs](#logs)
* [Transactions](#transactions)
* [Custom Queries](#custom-queries)

## Manual

#### Install

```bash
$ go get github.com/azer/crud
```

#### Initialize and Ping Database

````go
import (
  "github.com/azer/crud"
)

var DB crud.DB

func init () {
  var err error
  DB, err = crud.Open("mysql", os.Getenv("DATABASE_URL"))

  if err != nil {
    panic(err)
  }

  err := DB.Ping()
  if err != nil {
    panic(err)
  }
}
````

#### Define

```go
type User struct {
  Id int `sql:"auto-increment primary-key"`
  FirstName string
  LastName string
  ProfileId int
}

type Profile struct {
  Id int `sql:"auto-increment primary-key"`
  Bio string
}
```

#### Setup

`setup` makes sure the given structs have corresponding tables in the database.

```go
err := DB.Setup(User{}, Profile{})
```

#### Create

```go
user := &User{1, "Foo", "Bar", 1}
err := DB.Create(user)
```

#### Read

**Reading a single row:**

```go
var user *User
err := DB.Read(user, "WHERE id = ?", 1)
// => SELECT * FROM users WHERE id = 1

fmt.Println(user.Name)
// => Foo
```

**Reading multiple rows:**

```go
var users []*User

err := DB.Read(&users)
// => SELECT * FROM users

fmt.Println(len(users))
// => 10
```

**Custom queries:**

```go
var names []string{}
err := DB.Read(&names, "SELECT name FROM users")
```

```go
var totalUsers int
err := DB.Read(&totalUsers, "SELECT COUNT(id) FROM users"
```

#### Update

```go
var user *User
err := DB.Read(user, "WHERE id = ?", 1)

user.Name = "Yolo"
err := DB.Update(user)
```

#### Delete

````go
var user *User
err := DB.Read(user, "WHERE id = ?", 1)

user.Name = "Yolo"
err := DB.Delete(user)
````

#### Tables

To a create table:

```go
err := DB.CreateTable(User)
// => CREATE TABLE IF NOT EXISTS users ...
```

To drop a table:

```go
err := DB.DropTable(User)
// => DROP TABLE IF EXISTS users
```

#### Enabling Logs

If you want to see `crud` logs, specify `crud` in the `LOG` environment variable when you run your app. For example;

```
$ LOG=crud go run myapp.go
```

#### Transactions

```go
trans, err := DB.Begin()

trans.Create()
trans.Update()
trans.Delete()

err := trans.Commit()
```

#### Custom Queries

````go
result, err := DB.Query("DROP DATABASE yolo")
````
