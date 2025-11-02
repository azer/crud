package crud

import (
	stdsql "database/sql"
	"errors"
	"fmt"

	"github.com/azer/crud/v2/pg"
	"github.com/azer/crud/v2/sql"
)

func deleteRow(driver string, exec ExecFn, record interface{}) (stdsql.Result, error) {
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
		query = pg.DeleteQuery(table.SQLName, pk.SQL.Name)
	} else {
		query = sql.DeleteQuery(table.SQLName, pk.SQL.Name)
	}

	return exec(query, pk.Value)
}

func mustDelete(driver string, exec ExecFn, record interface{}) error {
	result, err := deleteRow(driver, exec, record)
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
