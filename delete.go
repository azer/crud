package crud

import (
	stdsql "database/sql"
	"errors"
	"fmt"

	"github.com/azer/crud/v2/sql"
)

func deleteRow(exec ExecFn, record interface{}) (stdsql.Result, error) {
	table, err := NewTable(record)

	if err != nil {
		return nil, err
	}

	pk := table.PrimaryKeyField()
	if pk == nil {
		return nil, errors.New(fmt.Sprintf("Table '%s' (%s) doesn't have a primary-key field", table.Name, table.SQLName))
	}

	return exec(sql.DeleteQuery(table.SQLName, pk.SQL.Name), pk.Value)
}

func mustDelete(exec ExecFn, record interface{}) error {
	result, err := deleteRow(exec, record)
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
