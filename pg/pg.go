package pg

import (
	"fmt"
	"strings"

	"github.com/azer/crud/v4/types"
)

func IsPostgres(driver string) bool {
	return driver == "postgres" || driver == "postgresql"
}

func NewTableQuery(name string, fields []*types.ColumnOptions, ifNotExists bool) string {
	ifNotExistsExt := ""
	if ifNotExists {
		ifNotExistsExt = " IF NOT EXISTS"
	}

	return fmt.Sprintf("CREATE TABLE%s %s (\n%s%s\n);",
		ifNotExistsExt, name, NewFieldQueries(fields), NewPrimaryKeyQuery(fields))
}

func NewFieldQueries(fields []*types.ColumnOptions) string {
	queries := []string{}
	for _, f := range fields {
		if f.Ignore {
			continue
		}
		queries = append(queries, NewFieldQuery(f))
	}
	return strings.Join(queries, ",\n")
}

func NewFieldQuery(field *types.ColumnOptions) string {
	if field.IsAutoIncrementing {
		return fmt.Sprintf("  %s SERIAL", field.Name)
	}

	length := ""
	required := ""
	defaultValue := ""
	unique := ""

	if field.Length > 0 {
		length = fmt.Sprintf("(%d)", field.Length)
	}

	if field.IsRequired {
		required = " NOT NULL"
	}

	if field.CurrentTimestamp {
		defaultValue = " DEFAULT CURRENT_TIMESTAMP"
	} else if field.DefaultValue != "" {
		defaultValue = fmt.Sprintf(" DEFAULT %s", field.DefaultValue)
	}

	if field.IsUnique {
		unique = " UNIQUE"
	}

	query := fmt.Sprintf("%s%s%s%s",
		length, required, defaultValue, unique)

	return fmt.Sprintf("  %s %s%s", field.Name, field.Type, query)
}

func NewPrimaryKeyQuery(fields []*types.ColumnOptions) string {
	keys := []string{}
	for _, f := range fields {
		if f.IsPrimaryKey {
			keys = append(keys, f.Name)
		}
	}

	if len(keys) == 0 {
		return ""
	}

	return fmt.Sprintf(",\n  PRIMARY KEY (%s)", strings.Join(keys, ", "))
}

func DropTableQuery(name string, ifExists bool) string {
	ext := ""
	if ifExists {
		ext = " IF EXISTS"
	}

	return fmt.Sprintf("DROP TABLE%s %s", ext, name)
}

func ShowTablesQuery(name string) string {
	return fmt.Sprintf("SELECT tablename FROM pg_tables WHERE schemaname = 'public' AND tablename = '%s'", name)
}

func DeleteQuery(tableName, index string) string {
	return fmt.Sprintf("DELETE FROM %s WHERE %s = $1", tableName, index)
}

func UpdateQuery(tableName, index string, columnNames []string) string {
	setParts := make([]string, len(columnNames))
	for i, col := range columnNames {
		setParts[i] = fmt.Sprintf("%s = $%d", col, i+1)
	}

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s = $%d",
		tableName, strings.Join(setParts, ", "), index, len(columnNames)+1)
}

func InsertQuery(tableName string, columnNames []string, autoIncrementField string) string {
	questionMarks := make([]string, len(columnNames))
	for i := range questionMarks {
		questionMarks[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName, strings.Join(columnNames, ", "), strings.Join(questionMarks, ", "))

	if autoIncrementField != "" {
		query += fmt.Sprintf(" RETURNING %s", autoIncrementField)
	}

	return query
}

func ConvertQueryPlaceholders(query string) string {
	// Convert ? placeholders to $1, $2, $3, etc.
	result := strings.Builder{}
	paramIndex := 1

	for i := 0; i < len(query); i++ {
		if query[i] == '?' {
			result.WriteString(fmt.Sprintf("$%d", paramIndex))
			paramIndex++
		} else {
			result.WriteByte(query[i])
		}
	}

	return result.String()
}
