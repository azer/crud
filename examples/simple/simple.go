package main

import (
	"fmt"
	"github.com/azer/crud"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type User struct {
	Id        int `sql:"auto-increment primary-key"`
	FirstName string
	LastName  string
}

func main() {
	DB, err := crud.Connect("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	if err := DB.CreateTables(User{}); err != nil {
		panic(err)
	}

	azer := User{
		FirstName: "Azer",
		LastName:  "Koculu",
	}

	if err := DB.Create(&azer); err != nil {
		panic(err)
	}

	copy := User{}
	if err := DB.Read(&copy, "WHERE first_name='Azer'"); err != nil {
		panic(err)
	}

	fmt.Println(copy.Id)
}
