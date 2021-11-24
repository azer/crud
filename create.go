package crud

import (
	stdsql "database/sql"
	"fmt"

	"github.com/azer/crud/v2/sql"
)

func createAndGetResult(exec ExecFn, record interface{}) (stdsql.Result, error) {
	row, err := newRow(record)
	if err != nil {
		return nil, err
	}

	columns := []string{}
	values := []interface{}{}

	for c, v := range row.SQLValues() {
		columns = append(columns, c)
		values = append(values, v)
	}

	return exec(sql.InsertQuery(row.SQLTableName, columns), values...)
}

func create(exec ExecFn, record interface{}) error {
	_, err := createAndGetResult(exec, record)
	return err
}

func createAndRead(exec ExecFn, query QueryFn, record interface{}) error {
	result, err := createAndGetResult(exec, record)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	table, err := NewTable(record)
	if err != nil {
		// this is a bad design choice made assuming that it'll never happen.
		return err
	}

	params := []interface{}{
		fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", table.SQLName, table.PrimaryKeyField().SQL.Name),
		id,
	}

	return read(query, record, params)
}
