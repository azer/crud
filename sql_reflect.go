package crud

import (
	"github.com/azer/crud/meta"
	"github.com/azer/snakecase"
	"github.com/jinzhu/inflection"
)

func SQLTableNameOf(any interface{}) string {
	if customTableName, ok := LookupCustomTableName(any); ok {
		return customTableName
	}

	return snakecase.SnakeCase(inflection.Plural(meta.TypeNameOf(any)))
}

func LookupCustomTableName(any interface{}) (string, bool) {
	if meta.IsSlice(any) {
		any = meta.CreateElement(any).Interface()
	}

	fields, err := GetFieldsOf(any)
	if err != nil {
		return "", false
	}

	for _, f := range fields {
		if len(f.SQL.TableName) > 0 {
			return f.SQL.TableName, true
		}
	}

	return "", false
}
