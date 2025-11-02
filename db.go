package crud

import (
	"context"
	stdsql "database/sql"

	"github.com/azer/crud/v3/pg"
	"github.com/azer/crud/v3/sql"
	"github.com/azer/logger"
	"github.com/labstack/gommon/random"
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

// Run any query on the database client, passing parameters optionally. Returns sql.Result.
func (db *DB) Exec(sql string, params ...interface{}) (stdsql.Result, error) {
	timer := log.Timer()
	result, error := db.Client.Exec(sql, params...)
	timer.End("SQL Query Executed: %s", sql)
	return result, error
}

// Run any query on the database client, passing parameters optionally. Its difference with
// `Exec` method is returning `sql.Rows` instead of `sql.Result`.
func (db *DB) Query(sql string, params ...interface{}) (*stdsql.Rows, error) {
	timer := log.Timer()
	result, error := db.Client.Query(sql, params...)
	timer.End("SQL Query Executed: %s", sql)
	return result, error
}

func (db *DB) CreateTable(st interface{}, ifexists bool) error {
	t, err := NewTable(db.Driver, st)
	if err != nil {
		return err
	}

	var query string
	if pg.IsPostgres(db.Driver) {
		query = pg.NewTableQuery(t.SQLName, t.SQLOptions(), ifexists)
	} else {
		query = sql.NewTableQuery(t.SQLName, t.SQLOptions(), ifexists)
	}

	_, err = db.Exec(query)
	return err
}

func (db *DB) DropTable(st interface{}, ifexists bool) error {
	t, err := NewTable(db.Driver, st)
	if err != nil {
		return err
	}

	var query string
	if pg.IsPostgres(db.Driver) {
		query = pg.DropTableQuery(t.SQLName, ifexists)
	} else {
		query = sql.DropTableQuery(t.SQLName, true)
	}

	_, err = db.Exec(query)

	return err
}

func (db *DB) CheckIfTableExists(name string) bool {
	var query string
	if pg.IsPostgres(db.Driver) {
		query = pg.ShowTablesQuery(name)
	} else {
		query = sql.ShowTablesLikeQuery(name)
	}

	var result string
	err := db.Client.QueryRow(query).Scan(&result)
	return err == nil && result == name
}

// Creates multiple tables from given any amount of structs. Calls `CreateTable` internally.
func (db *DB) CreateTables(structs ...interface{}) error {
	for _, st := range structs {
		if err := db.CreateTable(st, true); err != nil {
			return err
		}
	}

	return nil
}

// Drops correspoinding SQL tables of the given structs.
func (db *DB) DropTables(structs ...interface{}) error {
	for _, st := range structs {
		if err := db.DropTable(st, true); err != nil {
			return err
		}
	}

	return nil
}

// Drops (if they exist) and re-creates corresponding SQL tables for the given structs.
func (db *DB) ResetTables(structs ...interface{}) error {
	if err := db.DropTables(structs...); err != nil {
		return err
	}

	if err := db.CreateTables(structs...); err != nil {
		return err
	}

	return nil
}

func (db *DB) Create(record interface{}) error {
	return create(db.Driver, db.Exec, record)
}

func (db *DB) CreateAndRead(record interface{}) error {
	return createAndRead(db.Driver, db.Exec, db.Query, record)
}

func (db *DB) CreateAndGetResult(record interface{}) (stdsql.Result, error) {
	return createAndGetResult(db.Driver, db.Exec, record)
}

// Runs given SQL query and scans the result rows into the given target interface. The target
// interface could be both a single record or a slice of records.
//
// Usage Example:
//
// user := &User{}
// err := tx.Read(user, "SELECT * FROM users WHERE id = ?", 1)
//
// users := &[]*User{}
// err := tx.Read(users, "SELECT * FROM users", 1)
func (db *DB) Read(scanTo interface{}, params ...interface{}) error {
	return read(db.Driver, db.Query, scanTo, params)
}

// Generates and executes a DELETE query for given struct record. It matches the database row by finding
// out the primary key field defined in the table schema.
func (db *DB) Delete(record interface{}) error {
	return mustDelete(db.Driver, db.Exec, record)
}

// Start a DB transaction. It returns an interface w/ most of the methods DB provides.
func (db *DB) Begin(ctx context.Context) (*Tx, error) {
	client, err := db.Client.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{
		Id:      random.String(32),
		IdKey:   "Id",
		Driver:  db.Driver,
		Client:  client,
		Context: ctx,
	}, nil
}

// Return a database client that wraps underlying SQL execution methods with the context specified
func (db *DB) WithContext(ctx context.Context) *WithContext {
	return &WithContext{
		Context: ctx,
		DB:      db.Client,
		Driver:  db.Driver,
		Id:      random.String(32),
		IdKey:   "Id",
	}
}

// Establish DB connection and return a crud.DB instance w/ methods needed for accessing / writing the database.
// Example call: Connect("mysql", "root:123456@tcp(localhost:3306)/database_name?parseTime=true")
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

// Finding out the primary-key field of the given row, updates the corresponding record on the table
// with the values in the given record.
func (db *DB) Update(record interface{}) error {
	return mustUpdate(db.Driver, db.Exec, record)
}
