package crud

import (
	"fmt"
	"github.com/azer/crud/sql"
	"github.com/jmoiron/sqlx"
)

func Connect(driver, url string) (*DB, error) {
	client, err := sqlx.Connect(driver, url)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: client,
		Driver: driver,
		URL:    url,
	}, nil
}

type DB struct {
	Client *sqlx.DB
	Driver string
	URL    string
}

func (db *DB) Ping() error {
	return db.Client.Ping()
}

func (db *DB) ExecuteSQL(sql string, params ...interface{}) error {
	fmt.Println("$", sql)
	_, err := db.Client.Exec(sql, params...)
	return err
}

func (db *DB) CreateTable(structs ...interface{}) error {
	for _, st := range structs {
		t, err := NewTable(st)
		if err != nil {
			return err
		}

		if err := db.ExecuteSQL(sql.NewTableQuery(t.SQLName, t.SQLOptions(), true)); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) DropTable(structs ...interface{}) error {
	for _, st := range structs {
		t, err := NewTable(st)
		if err != nil {
			return err
		}

		if err := db.ExecuteSQL(sql.DropTableQuery(t.SQLName, true)); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) CheckIfTableExists(name string) bool {
	var result string
	err := db.Client.QueryRow(sql.ShowTablesLikeQuery(name)).Scan(&result)
	return err == nil && result == name
}

func (db *DB) Create(value interface{}) error {
	row, err := NewRow(value)
	if err != nil {
		return err
	}

	columns := []string{}
	values := []interface{}{}

	for c, v := range row.SQLValues() {
		columns = append(columns, c)
		values = append(values, v)
	}

	return db.ExecuteSQL(sql.InsertQuery(row.SQLTableName, columns), values...)
}
