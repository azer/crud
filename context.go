package crud

import (
	"context"
	stdsql "database/sql"

	"github.com/azer/logger"
)

type WithContext struct {
	Context context.Context
	DB      *stdsql.DB
	Id      string
	IdKey   string
}

// Execute any SQL query on the context client. Returns sql.Result.
func (ctx *WithContext) Exec(sql string, params ...interface{}) (stdsql.Result, error) {
	timer := log.Timer()
	result, err := ctx.DB.ExecContext(ctx.Context, sql, params...)
	timer.End("Executed SQL query.", logger.Attrs{
		ctx.IdKey: ctx.Id,
		"sql":     sql,
	})

	return result, err
}

// Execute any SQL query on the context client. Returns sql.Rows.
func (ctx *WithContext) Query(sql string, params ...interface{}) (*stdsql.Rows, error) {
	timer := log.Timer()
	result, err := ctx.DB.QueryContext(ctx.Context, sql, params...)
	timer.End("Run SQL query.", logger.Attrs{
		ctx.IdKey: ctx.Id,
		"sql":     sql,
	})

	return result, err
}

// Insert given record to the database.
func (ctx *WithContext) Create(record interface{}) error {
	return create(ctx.Exec, record)
}

// Inserts given record and scans the inserted row back to the given row.
func (ctx *WithContext) CreateAndRead(record interface{}) error {
	return createAndRead(ctx.Exec, ctx.Query, record)
}

// Run a select query on the databaase (w/ given parameters optionally) and scan the result(s) to the
// target interface specified as the first parameter.
func (ctx *WithContext) Read(scanTo interface{}, params ...interface{}) error {
	return read(ctx.Query, scanTo, params)
}

// Run an update query on the transaction, finding out the primary-key field of the given row.
func (ctx *WithContext) Update(record interface{}) error {
	return mustUpdate(ctx.Exec, record)
}

// Executes a DELETE query on the transaction for given struct record. It matches
// the database row by finding out the primary key field defined in the table schema.
func (ctx *WithContext) Delete(record interface{}) error {
	return mustDelete(ctx.Exec, record)
}
