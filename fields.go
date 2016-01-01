package crud

import (
	"github.com/azer/crud/sql"
)

type Field struct {
	Name  string
	Value interface{}
	SQL   *sql.Options
}

func GetFieldsOf(st interface{}) ([]*Field, error) {
	fields := []*Field{}

	iter := NewIteration(st)
	for iter.Next() {
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
