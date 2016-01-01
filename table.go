package crud

import (
	"github.com/azer/crud/sql"
	"github.com/azer/snakecase"
	"reflect"
	"strings"
)

func NewTable(st interface{}) (*Table, error) {
	fields, err := GetFieldsOf(st)
	if err != nil {
		return nil, err
	}

	name := GetTableNameOf(st)

	return &Table{
		Name:    name,
		SQLName: snakecase.SnakeCase(name),
		Fields:  fields,
	}, nil
}

type Table struct {
	Name    string
	SQLName string
	Fields  []*Field
}

func (table *Table) SQLOptions() []*sql.Options {
	result := []*sql.Options{}

	for _, f := range table.Fields {
		result = append(result, f.SQL)
	}

	return result
}

func GetTableNameOf(st interface{}) string {
	parts := strings.Split(reflect.TypeOf(st).String(), ".")
	return parts[len(parts)-1]
}
