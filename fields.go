package crud

import (
	"strings"

	"github.com/azer/crud/v2/types"
)

type Field struct {
	Name  string
	Value interface{}
	SQL   *types.ColumnOptions
}

// Get DB fields of any valid struct given
func GetFieldsOf(driver string, st interface{}) ([]*Field, error) {
	fields, err := collectFields(driver, st, []*Field{})
	if err != nil {
		return nil, err
	}

	return fields, nil
}

func collectFields(driver string, st interface{}, fields []*Field) ([]*Field, error) {
	iter := NewFieldIteration(driver, st)
	for iter.Next() {
		if iter.IsEmbeddedStruct() {
			if _fields, err := collectFields(driver, iter.ValueField().Interface(), fields); err != nil {
				return nil, err
			} else {
				fields = _fields
			}
			continue
		}

		sqlOptions, err := iter.SQLOptions()

		if err != nil {
			return nil, err
		}

		if sqlOptions.Ignore {
			continue
		}

		fields = append(fields, &Field{
			Name:  iter.Name(),
			Value: iter.Value(),
			SQL:   sqlOptions,
		})
	}

	return fields, nil
}

// If no PK is specified, then set `id` to be PK.
func SetDefaultPK(fields []*Field) {
	if HasPK(fields) {
		return
	}

	for i, f := range fields {
		sqlType := strings.ToLower(f.SQL.Type)
		isInt := sqlType == "int" || sqlType == "integer" || sqlType == "bigint"

		if !f.SQL.IsPrimaryKey && f.SQL.Name == "id" && isInt {
			fields[i].SQL.IsAutoIncrementing = true
			fields[i].SQL.AutoIncrement = 1
			fields[i].SQL.IsRequired = true
			fields[i].SQL.IsPrimaryKey = true
			return
		}
	}
}

func HasPK(fields []*Field) bool {
	for _, f := range fields {
		if f.SQL.IsPrimaryKey {
			return true
		}
	}

	return false
}
