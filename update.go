package crud

import (
	"errors"
	"fmt"
	"github.com/azer/crud/sql"
)

func (db *DB) Update(record interface{}) error {
	table, err := NewTable(record)
	if err != nil {
		return err
	}

	pk := table.PrimaryKeyField()
	if pk == nil {
		return errors.New(fmt.Sprintf("Table '%s' (%s) doesn't have a primary-key field", table.Name, table.SQLName))
	}

	_, err = db.Client.Exec(sql.UpdateQuery(table.SQLName, pk.SQL.Name, table.SQLUpdateColumnSet()), table.SQLUpdateValueSet()...)
	return err
}
