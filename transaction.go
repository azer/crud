package crud

import (
	"context"
	stdsql "database/sql"

	"github.com/azer/logger"
)

type Tx struct {
	Context context.Context
	Client  *stdsql.Tx
	Id      string
	IdKey   string
}

// Execute any SQL query on the transaction client. Returns sql.Result.
func (tx *Tx) Exec(sql string, params ...interface{}) (stdsql.Result, error) {
	timer := log.Timer()
	result, err := tx.Client.ExecContext(tx.Context, sql, params...)
	timer.End("Executed SQL query.", logger.Attrs{
		tx.IdKey: tx.Id,
		"sql":    sql,
	})
	return result, err
}

// Execute any SQL query on the transaction client. Returns sql.Rows.
func (tx *Tx) Query(sql string, params ...interface{}) (*stdsql.Rows, error) {
	timer := log.Timer()
	result, err := tx.Client.QueryContext(tx.Context, sql, params...)
	timer.End("Run SQL query.", logger.Attrs{
		tx.IdKey: tx.Id,
		"sql":    sql,
	})
	return result, err
}

// Commit the transaction.
func (tx *Tx) Commit() error {
	log.Info("Committing", logger.Attrs{
		tx.IdKey: tx.Id,
	})
	return tx.Client.Commit()
}

// Rollback the transaction.
func (tx *Tx) Rollback() error {
	log.Info("Rolling back", logger.Attrs{
		tx.IdKey: tx.Id,
	})
	return tx.Client.Rollback()
}

// Insert given record to the database.
func (tx *Tx) Create(record interface{}) error {
	return create(tx.Exec, record)
}

// Inserts given record and scans the inserted row back to the given row.
func (tx *Tx) CreateAndRead(record interface{}) error {
	return createAndRead(tx.Exec, tx.Query, record)
}

// Run a select query on the databaase (w/ given parameters optionally) and scan the result(s) to the
// target interface specified as the first parameter.
//
// Usage Example:
//
// user := &User{}
// err := tx.Read(user, "SELECT * FROM users WHERE id = ?", 1)
//
// users := &[]*User{}
// err := tx.Read(users, "SELECT * FROM users", 1)
//
func (tx *Tx) Read(scanTo interface{}, params ...interface{}) error {
	return read(tx.Query, scanTo, params)
}

// Run an update query on the transaction, finding out the primary-key field of the given row.
func (tx *Tx) Update(record interface{}) error {
	return mustUpdate(tx.Exec, record)
}

// Executes a DELETE query on the transaction for given struct record. It matches
// the database row by finding out the primary key field defined in the table schema.
func (tx *Tx) Delete(record interface{}) error {
	return mustDelete(tx.Exec, record)
}
