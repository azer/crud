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

	_, err = db.Exec(sql.InsertQuery(row.SQLTableName, columns), values...)
	return err
}
