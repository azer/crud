package crud

import (
	stdsql "database/sql"
	"fmt"
	"github.com/azer/crud/sql"
)

func CreateAndGetResult(exec ExecFn, record interface{}) (stdsql.Result, error) {
	row, err := NewRow(record)
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

func Create(exec ExecFn, record interface{}) error {
	_, err := CreateAndGetResult(exec, record)
	return err
}

func CreateAndRead(exec ExecFn, query QueryFn, record interface{}) error {
	result, err := CreateAndGetResult(exec, record)
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
		fmt.Sprintf("WHERE %s = ?", table.PrimaryKeyField().SQL.Name),
		id,
	}

	return Read(query, record, params)
}
