package crud

import (
	"context"
	stdsql "database/sql"

	"github.com/azer/crud/v2/sql"
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

// Takes any valid struct and creates a SQL table from it.
func (db *DB) CreateTable(st interface{}, ifexists bool) error {
	t, err := NewTable(st)
	if err != nil {
		return err
	}

	_, err = db.Exec(sql.NewTableQuery(t.SQLName, t.SQLOptions(), ifexists))
	return err
}

// Takes any valid struct, finds out its corresponding SQL table and drops it.
func (db *DB) DropTable(st interface{}, ifexists bool) error {
	t, err := NewTable(st)
	if err != nil {
		return err
	}

	_, err = db.Exec(sql.DropTableQuery(t.SQLName, true))
	return err
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

// Runs a query to check if the given table exists and returns bool
func (db *DB) CheckIfTableExists(name string) bool {
	var result string
	err := db.Client.QueryRow(sql.ShowTablesLikeQuery(name)).Scan(&result)
	return err == nil && result == name
}

// Inserts given record into the database, generating an insert query for it.
func (db *DB) Create(record interface{}) error {
	return create(db.Exec, record)
}

func (db *DB) CreateAndGetResult(record interface{}) (stdsql.Result, error) {
	return createAndGetResult(db.Exec, record)
}

// Inserts given record and scans the inserted row back to the given row.
func (db *DB) CreateAndRead(record interface{}) error {
	return createAndRead(db.Exec, db.Query, record)
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
//
func (db *DB) Read(scanTo interface{}, params ...interface{}) error {
	return read(db.Query, scanTo, params)
}

// Finding out the primary-key field of the given row, updates the corresponding record on the table
// with the values in the given record.
func (db *DB) Update(record interface{}) error {
	return mustUpdate(db.Exec, record)
}

// Generates and executes a DELETE query for given struct record. It matches the database row by finding
// out the primary key field defined in the table schema.
func (db *DB) Delete(record interface{}) error {
	return mustDelete(db.Exec, record)
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
		Client:  client,
		Context: ctx,
	}, nil
}

// Return a database client that wraps underlying SQL execution methods with the context specified
func (db *DB) WithContext(ctx context.Context) *WithContext {
	return &WithContext{
		Context: ctx,
		DB:      db.Client,
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
