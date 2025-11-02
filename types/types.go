package types

import "strings"

var ValidMySQLTypes = map[string]int{
	"int":       11,
	"bigint":    20,
	"varchar":   255,
	"text":      -1,
	"date":      -1,
	"time":      -1,
	"timestamp": -1,
}

var ValidPostgresTypes = map[string]int{
	"integer":          -1,
	"bigint":           -1,
	"varchar":          255,
	"text":             -1,
	"timestamp":        -1,
	"boolean":          -1,
	"bytea":            -1,
	"double precision": -1,
}

// IsValidSqlType checks if a SQL type is valid for either MySQL or PostgreSQL.
// The input is converted to lowercase before checking.
func IsValidSqlType(sqlType string) bool {
	lowerType := strings.ToLower(sqlType)

	_, mysqlValid := ValidMySQLTypes[lowerType]
	_, pgValid := ValidPostgresTypes[lowerType]

	return mysqlValid || pgValid
}

func GetValidSqlType(sqlType string) (length int, ok bool) {
	lowerType := strings.ToLower(sqlType)

	// Check MySQL types first
	if length, ok := ValidMySQLTypes[lowerType]; ok {
		return length, true
	}

	// Check PostgreSQL types
	if length, ok := ValidPostgresTypes[lowerType]; ok {
		return length, true
	}

	return -1, false
}
