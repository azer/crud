package crud

import (
	"github.com/azer/crud/reflect"
	"github.com/azer/crud/sql"
	"github.com/azer/snakecase"
)

func NewTable(any interface{}) (*Table, error) {
	if reflect.IsSlice(any) {
		any = reflect.CreateElement(any).Interface()
	}

	fields, err := GetFieldsOf(any)
	if err != nil {
		return nil, err
	}

	name := reflect.TypeNameOf(any)

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

func (table *Table) SQLColumnDict() map[string]string {
	result := map[string]string{}

	for _, field := range table.Fields {
		result[field.SQL.Name] = field.Name
	}

	return result
}
