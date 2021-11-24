package crud

import (
	"github.com/azer/crud/v2/meta"
	"github.com/azer/snakecase"
	"github.com/jinzhu/inflection"
)

// Find out what given interface should be called in the database. It first looks up
// if a table name was explicitly specified (see "table-name" option), or automatically
// generates a plural name from the name of the struct type.
func SQLTableNameOf(any interface{}) string {
	if customTableName, ok := lookupCustomTableName(any); ok {
		return customTableName
	}

	return snakecase.SnakeCase(inflection.Plural(meta.TypeNameOf(any)))
}

func lookupCustomTableName(any interface{}) (string, bool) {
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
