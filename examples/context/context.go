package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/azer/crud/v2"
	_ "github.com/go-sql-driver/mysql"
)

var DB *crud.DB

type User struct {
	Id        int `sql:"auto-increment primary-key"`
	FirstName string
	LastName  string
}

func main() {
	var err error
	DB, err = crud.Connect("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	if err := DB.CreateTables(User{}); err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	db := DB.WithContext(ctx)

	row := User{
		FirstName: "Foo",
		LastName:  "Bar",
	}

	if err := db.CreateAndRead(&row); err != nil {
		panic(err)
	}

	/*if _, err := db.Query("SELECT sleep(15)"); err != nil {
		panic(err)
	}*/

	w.Write([]byte(fmt.Sprintf("%d / %s / %s", row.Id, row.FirstName, row.LastName)))
}
