package crud

import (
	stdsql "database/sql"
	"errors"
	"fmt"

	"github.com/azer/crud/v4/pg"
	"github.com/azer/crud/v4/sql"
)

func update(driver string, exec ExecFn, record interface{}) (stdsql.Result, error) {
	table, err := NewTable(driver, record)
	if err != nil {
		return nil, err
	}

	pk := table.PrimaryKeyField()
	if pk == nil {
		return nil, errors.New(fmt.Sprintf("Table '%s' (%s) doesn't have a primary-key field", table.Name, table.SQLName))
	}

	var query string
	if pg.IsPostgres(driver) {
		query = pg.UpdateQuery(table.SQLName, pk.SQL.Name, table.SQLUpdateColumnSet())
	} else {
		query = sql.UpdateQuery(table.SQLName, pk.SQL.Name, table.SQLUpdateColumnSet())
	}

	return exec(query, table.SQLUpdateValueSet()...)
}

func mustUpdate(driver string, exec ExecFn, record interface{}) error {
	result, err := update(driver, exec, record)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return stdsql.ErrNoRows
	}

	return nil
}
