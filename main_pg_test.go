//go:build postgres
// +build postgres

package crud_test

import (
	"fmt"
	"os"

	"database/sql"
	"time"

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

func getDriver() string {
	return "postgres"
}

func getIntType() string {
	return "INTEGER"
}

func getPkType() string {
	return "SERIAL"
}

func getBigintType() string {
	return "BIGINT"
}

func getVarcharType() string {
	return "VARCHAR"
}

func getTextType() string {
	return "TEXT"
}

func getBlobType() string {
	return "BYTEA"
}

func getBoolType() string {
	return "BOOLEAN"
}

var DB *crud.DB

type UserProfile struct {
	Id         int    `json:"id" sql:"auto-increment primary-key required"`
	Name       string `json:"name" sql:"required"`
	Bio        string `json:"bio" sql:"type=text"`
	Email      string `json:"e-mail" sql:"name=email unique"`
	Attachment []byte `json:"attachment"`
	Modified   int64  `json:"modified" sql:"name=modified_col"`
}

type UserProfileNull struct {
	Id       sql.NullInt64  `json:"id" sql:"auto-increment primary-key required"`
	Name     sql.NullString `json:"name" sql:"required"`
	Bio      sql.NullString `json:"bio" sql:"type=text"`
	Email    sql.NullString `json:"e-mail" sql:"name=email"`
	Modified sql.NullInt64  `json:"modified" sql:"name=modified"`
}

type Mixed struct {
	Id        int    `json:"-" sql:"  primary-key auto-increment unsigned name=id table-name=__mixed__ "`
	UserId    int    `json:"-" valid:"User.Id~Specified user was not found" sql:" name=user_id"`
	Secret    string `json:"-" valid:"required" sql:" name=secret"`
	CreatedAt int64  `json:"-" sql:"default=0 name=created_at"`
	UpdatedAt int64  `json:"-" sql:"default=0 name=updated_at"`
}

type Post struct {
	Id        int       `json:"id" sql:"auto-increment primary-key required table-name=renamed_posts"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type Foo struct {
	Id     int
	APIKey string
	YOLO   bool
	Beast  string
}

type EmbeddedFoo struct {
	Foo
	Span int
	Eggs string
}

type FooSlice []Foo
type FooPTRSlice []*Foo

type CustomTableName struct {
	Foo int `sql:"table-name=yolo"`
}
