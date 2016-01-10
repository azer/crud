package crud

import (
	stdsql "database/sql"
	"github.com/azer/crud/sql"
)

func Connect(driver, url string) (*DB, error) {
	client, err := stdsql.Open(driver, url)
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
	Client *stdsql.DB
	Driver string
	URL    string
}

func (db *DB) Ping() error {
	return db.Client.Ping()
}

func (db *DB) ExecuteSQL(sql string, params ...interface{}) error {
	_, err := db.Client.Exec(sql, params...)
	return err
}

func (db *DB) CreateTable(st interface{}, ifexists bool) error {
	t, err := NewTable(st)
	if err != nil {
		return err
	}

	return db.ExecuteSQL(sql.NewTableQuery(t.SQLName, t.SQLOptions(), ifexists))
}

func (db *DB) DropTable(st interface{}, ifexists bool) error {
	t, err := NewTable(st)
	if err != nil {
		return err
	}

	return db.ExecuteSQL(sql.DropTableQuery(t.SQLName, true))
}

func (db *DB) CreateTables(structs ...interface{}) error {
	for _, st := range structs {
		if err := db.CreateTable(st, true); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) DropTables(structs ...interface{}) error {
	for _, st := range structs {
		if err := db.DropTable(st, true); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) ResetTables(structs ...interface{}) error {
	if err := db.DropTables(structs...); err != nil {
		return err
	}

	if err := db.CreateTables(structs...); err != nil {
		return err
	}

	return nil
}

func (db *DB) CheckIfTableExists(name string) bool {
	var result string
	err := db.Client.QueryRow(sql.ShowTablesLikeQuery(name)).Scan(&result)
	return err == nil && result == name
}

func (db *DB) CompleteSelectQuery(query string, scanner *Scan) string {
	if scanner.Table == nil {
		return query
	}

	return sql.CompleteSelectQuery(scanner.Table.SQLName, []string{}, query)
}
