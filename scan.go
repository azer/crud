package crud

import (
	"database/sql"
	"reflect"

	"github.com/azer/crud/v2/meta"
)

// Create a scanner for any given interface. This function will be called for
// every target interface passed to DB methods that scans results.
func NewScan(to interface{}) (*Scan, error) {
	scan := &Scan{
		To:         to,
		ToPointers: meta.HasPointers(to),
		ToStructs:  meta.HasAnyStruct(to),
	}

	if scan.ToStructs {
		table, err := NewTable(to)
		if err != nil {
			return nil, err
		}

		scan.Table = table
		scan.SQLColumnDict = table.SQLColumnDict()
	}

	return scan, nil
}

type Scan struct {
	To            interface{}
	ToPointers    bool
	ToStructs     bool
	SQLColumnDict map[string]string
	Table         *Table
}

func (scan *Scan) All(rows *sql.Rows) error {
	writeTo := meta.Addressable(scan.To)

	for rows.Next() {
		record := meta.CreateElement(scan.To)

		if err := scan.Scan(rows, record); err != nil {
			return err
		}

		meta.Push(writeTo, record)
	}

	return nil
}

func (scan *Scan) One(rows *sql.Rows) error {
	for rows.Next() {
		return scan.Scan(rows, meta.DirectValueOf(scan.To))
	}

	return sql.ErrNoRows
}

func (scan *Scan) Scan(rows *sql.Rows, record reflect.Value) error {
	if scan.ToStructs {
		return scan.ScanToStruct(rows, record)
	}

	if record.Kind() != reflect.Ptr {
		return rows.Scan(record.Addr().Interface())
	} else {
		return rows.Scan(record.Interface())
	}
}

func (scan *Scan) ScanToStruct(rows *sql.Rows, record reflect.Value) error {
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))

	for i, column := range columns {
		var field reflect.Value

		fieldName := scan.SQLColumnDict[column]

		if scan.ToPointers {
			field = record.Elem().FieldByName(fieldName)
		} else {
			field = record.FieldByName(fieldName)
		}

		if field.IsValid() {
			values[i] = field.Addr().Interface()
		} else {
			values[i] = &values[i]
		}
	}

	return rows.Scan(values...)
}
