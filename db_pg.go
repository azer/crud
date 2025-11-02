package crud

import (
	"fmt"
	"strings"

	"github.com/azer/crud/v2/sql"
)

func isPostgres(driver string) bool {
	return driver == "postgres" || driver == "postgresql"
}

func postgresCreateTableQuery(name string, fields []*sql.Options, ifNotExists bool) string {
	ifNotExistsExt := ""
	if ifNotExists {
		ifNotExistsExt = " IF NOT EXISTS"
	}

	return fmt.Sprintf("CREATE TABLE%s %s (\n%s%s\n);",
		ifNotExistsExt, name, postgresFieldQueries(fields), postgresPrimaryKeyQuery(fields))
}

func postgresFieldQueries(fields []*sql.Options) string {
	queries := []string{}
	for _, f := range fields {
		if f.Ignore {
			continue
		}
		queries = append(queries, postgresFieldQuery(f))
	}
	return strings.Join(queries, ",\n")
}

func postgresFieldQuery(field *sql.Options) string {
	if field.IsAutoIncrementing {
		return fmt.Sprintf("  %s SERIAL", field.Name)
	}

	pgType := postgresType(field.Type, field.Length)
	var parts []string
	parts = append(parts, fmt.Sprintf("  %s %s", field.Name, pgType))

	if field.IsRequired {
		parts = append(parts, "NOT NULL")
	}

	if field.DefaultValue != "" {
		parts = append(parts, fmt.Sprintf("DEFAULT %s", field.DefaultValue))
	}

	if field.IsUnique {
		parts = append(parts, "UNIQUE")
	}

	return strings.Join(parts, " ")
}

func postgresType(sqlType string, length int) string {
	typeMap := map[string]string{
		"int":       "INTEGER",
		"bigint":    "BIGINT",
		"varchar":   fmt.Sprintf("VARCHAR(%d)", length),
		"text":      "TEXT",
		"date":      "DATE",
		"time":      "TIME",
		"timestamp": "TIMESTAMP",
		"tinyint":   "BOOLEAN",
		"blob":      "BYTEA",
		"float":     "DOUBLE PRECISION",
	}

	if pgType, ok := typeMap[sqlType]; ok {
		return pgType
	}
	return sqlType
}

func postgresPrimaryKeyQuery(fields []*sql.Options) string {
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

func postgresInsertQuery(tableName string, columnNames []string, autoIncrementField string) string {
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

func postgresUpdateQuery(tableName, index string, columnNames []string) string {
	setParts := make([]string, len(columnNames))
	for i, col := range columnNames {
		setParts[i] = fmt.Sprintf("%s = $%d", col, i+1)
	}

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s = $%d",
		tableName, strings.Join(setParts, ", "), index, len(columnNames)+1)
}

func postgresDeleteQuery(tableName, index string) string {
	return fmt.Sprintf("DELETE FROM %s WHERE %s = $1", tableName, index)
}

func postgresShowTablesQuery(name string) string {
	return fmt.Sprintf("SELECT tablename FROM pg_tables WHERE schemaname = 'public' AND tablename = '%s'", name)
}

func postgresDropTableQuery(name string, ifExists bool) string {
	ext := ""
	if ifExists {
		ext = " IF EXISTS"
	}
	return fmt.Sprintf("DROP TABLE%s %s", ext, name)
}
