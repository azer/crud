package crud

import (
	stdsql "database/sql"
	"github.com/azer/crud/sql"
	"github.com/azer/logger"
)

var log = logger.New("crud")

type ExecFn func(string, ...interface{}) (stdsql.Result, error)
type QueryFn func(string, ...interface{}) (*stdsql.Rows, error)

type DB struct {
	Client *stdsql.DB
	Driver string
	URL    string
}

func (db *DB) Ping() error {
	return db.Client.Ping()
}

func (db *DB) Exec(sql string, params ...interface{}) (stdsql.Result, error) {
	log.Info(sql)
	return db.Client.Exec(sql, params...)
}

func (db *DB) Query(sql string, params ...interface{}) (*stdsql.Rows, error) {
	log.Info(sql)
	return db.Client.Query(sql, params...)
}

func (db *DB) CreateTable(st interface{}, ifexists bool) error {
	t, err := NewTable(st)
	if err != nil {
		return err
	}

	_, err = db.Exec(sql.NewTableQuery(t.SQLName, t.SQLOptions(), ifexists))
	return err
}

func (db *DB) DropTable(st interface{}, ifexists bool) error {
	t, err := NewTable(st)
	if err != nil {
		return err
	}

	_, err = db.Exec(sql.DropTableQuery(t.SQLName, true))
	return err
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

func (db *DB) Create(record interface{}) error {
	return Create(db.Exec, record)
}

func (db *DB) Read(scanTo interface{}, params ...interface{}) error {
	return Read(db.Query, scanTo, params)
}

func (db *DB) Update(record interface{}) error {
	return MustUpdate(db.Exec, record)
}

func (db *DB) Delete(record interface{}) error {
	return MustDelete(db.Exec, record)
}

func (db *DB) Begin() (*Tx, error) {
	client, err := db.Client.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{
		Client: client,
	}, nil
}

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
