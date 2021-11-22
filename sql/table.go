package sql

import (
	"fmt"
	"strings"
)

func NewTableQuery(name string, fields []*Options, ifNotExists bool) string {
	ifNotExistsExt := ""
	if ifNotExists {
		ifNotExistsExt = " IF NOT EXISTS"
	}

	return fmt.Sprintf("CREATE TABLE%s `%s` (\n%s%s\n)%s;",
		ifNotExistsExt, name, NewFieldQueries(fields), NewPrimaryKeyQuery(fields), NewTableConfigQuery(fields))
}

func NewFieldQueries(fields []*Options) string {
	queries := []string{}

	for _, f := range fields {
		if f.Ignore {
			continue
		}

		queries = append(queries, NewFieldQuery(f))
	}

	return strings.Join(queries, ",\n")
}

func NewFieldQuery(field *Options) string {
	length := ""
	autoIncrement := ""
	required := ""
	defaultValue := ""
	unsigned := ""
	unique := ""

	if field.Length > -1 {
		length = fmt.Sprintf("(%d)", field.Length)
	}

	if field.AutoIncrement > 0 {
		autoIncrement = " AUTO_INCREMENT"
	}

	if field.IsRequired {
		required = " NOT NULL"
	}

	if field.DefaultValue != "" {
		defaultValue = fmt.Sprintf(" DEFAULT %s", field.DefaultValue)
	}

	if field.IsUnsigned {
		unsigned = " UNSIGNED"
	}

	if field.IsUnique {
		unique = " UNIQUE"
	}

	query := fmt.Sprintf("%s%s%s%s%s%s",
		length, required, defaultValue, unsigned, unique, autoIncrement)

	return fmt.Sprintf("  `%s` %s%s", field.Name, field.Type, query)
}

func NewPrimaryKeyQuery(fields []*Options) string {
	keys := []string{}

	for _, f := range fields {
		if f.IsPrimaryKey {
			keys = append(keys, f.Name)
		}
	}

	if len(keys) == 0 {
		return ""
	}

	return fmt.Sprintf(",\n  PRIMARY KEY (`%s`)", strings.Join(keys, "`, `"))
}

func NewTableConfigQuery(fields []*Options) string {
	autoIncrement := ""
	for _, f := range fields {
		if f.AutoIncrement > 1 {
			autoIncrement = fmt.Sprintf(" AUTO_INCREMENT=%d", f.AutoIncrement)
		}
	}

	return fmt.Sprintf("%s", autoIncrement)
}

func DropTableQuery(name string, ifExists bool) string {
	ext := ""

	if ifExists {
		ext = " IF EXISTS"
	}

	return fmt.Sprintf("DROP TABLE%s %s", ext, name)
}

func ShowTablesLikeQuery(name string) string {
	return fmt.Sprintf("SHOW TABLES LIKE '%s'", name)
}

func InsertQuery(tableName string, columnNames []string) string {
	var questionMarks string

	if len(columnNames) > 0 {
		questionMarks = strings.Repeat("?,", len(columnNames))
		questionMarks = questionMarks[:len(questionMarks)-1]
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName, strings.Join(quoteColumnNames(columnNames), ","), questionMarks)
}

func SelectQuery(tableName string, columnNames []string) string {
	columns := strings.Join(columnNames, ",")
	if columns == "" {
		columns = "*"
	}

	return fmt.Sprintf("SELECT %s FROM %s", columns, tableName)
}

func UpdateQuery(tableName, index string, columnNames []string) string {
	return fmt.Sprintf("%s WHERE %s=?", UpdateAllQuery(tableName, columnNames), index)
}

func UpdateAllQuery(tableName string, columnNames []string) string {
	return fmt.Sprintf("UPDATE %s SET %s=?", tableName, strings.Join(quoteColumnNames(columnNames), "=?, "))
}

func DeleteQuery(tableName, index string) string {
	return fmt.Sprintf("DELETE FROM %s WHERE %s=?", tableName, index)
}

func quoteColumnNames(columns []string) []string {
	quoted := []string{}

	for _, c := range columns {
		quoted = append(quoted, fmt.Sprintf("`%s`", c))
	}

	return quoted
}
