//go:build postgres
// +build postgres

package crud_test

import (
	"fmt"
	"os"

	"github.com/azer/crud/v2"
	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("db:", os.Getenv("DATABASE_URL"))

	var err error
	DB, err = crud.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
}
