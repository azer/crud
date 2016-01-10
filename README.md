## crud

A minimalistic database library for Go, with simple and familiar interface.

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

#### Initialize and Ping

```go
import (
  "github.com/azer/crud"
  _ "github.com/go-sql-driver/mysql"
)

var DB *crud.DB

func init () {
  DB, err := crud.Connect("mysql", os.Getenv("DATABASE_URL"))
  err := DB.Ping()
}
```

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
  Bio string `sql:"text"`
}
```

#### Create Tables

`CreateTables` takes list of structs and makes sure they exist in the database.

```go
err := DB.CreateTables(User{}, Profile{})
```

#### Create

```go
user := &User{1, "Foo", "Bar", 1}
err := DB.Create(user)
```

#### Read

**Reading a single row:**

```go
user := User{}
err := DB.Read(user, "WHERE id = ?", 1) // You can type the full query if preferred.
// => SELECT * FROM users WHERE id = 1

fmt.Println(user.Name)
// => Foo
```

**Reading multiple rows:**

```go
users := []*User{}

err := DB.Read(&users)
// => SELECT * FROM users

fmt.Println(len(users))
// => 10
```

**Scanning to custom values:**

```go
names := []string{}
err := DB.Read(&names, "SELECT name FROM users")
```

```
name := ""
err := DB.Read(&name, "SELECT name FROM users WHERE id=1")
```

```go
totalUsers := 0
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
