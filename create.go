package crud

import (
	stdsql "database/sql"
	"fmt"

	"github.com/azer/crud/v4/pg"
	"github.com/azer/crud/v4/sql"
)

func createAndGetResult(driver string, exec ExecFn, record interface{}) (stdsql.Result, error) {
	row, err := newRow(driver, record)
	if err != nil {
		return nil, err
	}

	columns := []string{}
	values := []interface{}{}

	for c, v := range row.SQLValues() {
		columns = append(columns, c)
		values = append(values, v)
	}

	if pg.IsPostgres(driver) {
		return exec(pg.InsertQuery(row.SQLTableName, columns, ""), values...)
	}

	return exec(sql.InsertQuery(row.SQLTableName, columns), values...)
}

func create(driver string, exec ExecFn, record interface{}) error {
	_, err := createAndGetResult(driver, exec, record)
	return err
}

func createAndRead(driver string, exec ExecFn, query QueryFn, record interface{}) error {
	table, err := NewTable(driver, record)
	if err != nil {
		return err
	}

	row, err := newRow(driver, record)
	if err != nil {
		return err
	}

	pkField := table.PrimaryKeyField()
	if pkField == nil {
		return fmt.Errorf("no primary key field found")
	}

	columns := []string{}
	values := []interface{}{}

	for c, v := range row.SQLValues() {
		// Skip auto-increment primary key for postgres
		if pg.IsPostgres(driver) && c == pkField.SQL.Name && pkField.SQL.IsAutoIncrementing {
			continue
		}

		columns = append(columns, c)
		values = append(values, v)
	}

	if pg.IsPostgres(driver) {
		queryStr := pg.InsertQuery(row.SQLTableName, columns, pkField.SQL.Name)
		rows, err := query(queryStr, values...)
		if err != nil {
			return err
		}

		defer rows.Close()

		if rows.Next() {
			var id int64
			if err := rows.Scan(&id); err != nil {
				return err
			}

			params := []interface{}{
				fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", table.SQLName, pkField.SQL.Name),
				id,
			}
			return read(driver, query, record, params)
		}

		return fmt.Errorf("no id returned")
	}

	result, err := exec(sql.InsertQuery(row.SQLTableName, columns), values...)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	params := []interface{}{
		fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", table.SQLName, pkField.SQL.Name),
		id,
	}

	return read(driver, query, record, params)
}
