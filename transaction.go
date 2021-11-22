package crud

import (
	stdsql "database/sql"
)

type Tx struct {
	Client *stdsql.Tx
}

// Execute any SQL query on the transaction client. Returns sql.Result.
func (tx *Tx) Exec(sql string, params ...interface{}) (stdsql.Result, error) {
	return tx.Client.Exec(sql, params...)
}

// Execute any SQL query on the transaction client. Returns sql.Rows.
func (tx *Tx) Query(sql string, params ...interface{}) (*stdsql.Rows, error) {
	return tx.Client.Query(sql, params...)
}

// Commit the transaction.
func (tx *Tx) Commit() error {
	return tx.Client.Commit()
}

// Rollback the transaction.
func (tx *Tx) Rollback() error {
	return tx.Client.Rollback()
}

// Insert given record to the database.
func (tx *Tx) Create(record interface{}) error {
	return Create(tx.Exec, record)
}

// Run a select query on the databaase (w/ given parameters optionally) and scan the result(s) to the
// target interface specified as the first parameter.
//
// Usage Example:
//
// user := &User{}
// err := tx.Read(user, "SELECT * FROM users WHERE id = ?", 1)
//
func (tx *Tx) Read(scanTo interface{}, params ...interface{}) error {
	return Read(tx.Query, scanTo, params)
}

// Run an update query on the transaction, finding out the primary-key field of the given row.
func (tx *Tx) Update(record interface{}) error {
	return MustUpdate(tx.Exec, record)
}

// Executes a DELETE query on the transaction for given struct record. It matches
// the database row by finding out the primary key field defined in the table schema.
func (tx *Tx) Delete(record interface{}) error {
	return MustDelete(tx.Exec, record)
}
