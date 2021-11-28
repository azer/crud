package crud

import (
	"github.com/azer/crud/v2/meta"
	"github.com/azer/crud/v2/sql"
)

// Create an internal representation of a database table, including its fields from given
// struct record
func NewTable(any interface{}) (*Table, error) {
	if meta.IsSlice(any) {
		any = meta.CreateElement(any).Interface()
	}

	fields, err := GetFieldsOf(any)
	if err != nil {
		return nil, err
	}

	SetDefaultPK(fields)

	name, sqlName := ReadTableName(any)

	return &Table{
		Name:    name,
		SQLName: sqlName,
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

func (table *Table) PrimaryKeyField() *Field {
	for _, f := range table.Fields {
		if f.SQL.IsPrimaryKey {
			return f
		}
	}

	return nil
}

func (table *Table) SQLUpdateColumnSet() []string {
	columns := []string{}

	for _, f := range table.Fields {
		if f.SQL.Ignore || f.SQL.IsAutoIncrementing {
			continue
		}

		columns = append(columns, f.SQL.Name)
	}

	return columns
}

func (table *Table) SQLUpdateValueSet() []interface{} {
	values := []interface{}{}

	for _, f := range table.Fields {
		if f.SQL.Ignore || f.SQL.IsAutoIncrementing {
			continue
		}

		values = append(values, f.Value)
	}

	pk := table.PrimaryKeyField()

	if pk != nil {
		values = append(values, pk.Value)
	}

	return values
}

// Return struct name and SQL table name
func ReadTableName(any interface{}) (string, string) {
	return meta.TypeNameOf(any), SQLTableNameOf(any)
}

// Return table columns for given struct, pointer to struct or slice of structs.
func ReadTableColumns(any interface{}) ([]string, error) {
	if meta.IsSlice(any) {
		any = meta.CreateElement(any).Interface()
	}

	fields, err := GetFieldsOf(any)
	if err != nil {
		return nil, err
	}

	columns := []string{}

	for _, col := range fields {
		columns = append(columns, col.SQL.Name)
	}

	return columns, nil
}
