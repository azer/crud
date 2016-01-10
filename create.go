package crud

import (
	"github.com/azer/crud/sql"
)

func (db *DB) Create(value interface{}) error {
	row, err := NewRow(value)
	if err != nil {
		return err
	}

	columns := []string{}
	values := []interface{}{}

	for c, v := range row.SQLValues() {
		columns = append(columns, c)
		values = append(values, v)
	}

	return db.ExecuteSQL(sql.InsertQuery(row.SQLTableName, columns), values...)
}
