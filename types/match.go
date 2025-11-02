package types

import (
	"errors"
	"fmt"
)

var TypeLength = map[string]map[string]int{
	"postgres": {
		"VARCHAR": 255,
	},
	"mysql": {
		"int":            11,
		"uint":           11,
		"int64":          20,
		"uint64":         20,
		"string":         255,
		"bigint":         20,
		"varchar":        255,
		"sql.NullInt64":  20,
		"sql.NullString": 255,
	},
}

var TypeDict = map[string]map[string]string{
	"postgres": {
		"float32":         "DOUBLE PRECISION",
		"float64":         "DOUBLE PRECISION",
		"int":             "INTEGER",
		"uint":            "INTEGER",
		"int64":           "BIGINT",
		"uint64":          "BIGINT",
		"string":          "VARCHAR",
		"time.Time":       "TIMESTAMP",
		"bool":            "BOOLEAN",
		"[]byte":          "BYTEA",
		"[]uint8":         "BYTEA",
		"sql.NullFloat32": "DOUBLE PRECISION",
		"sql.NullFloat64": "DOUBLE PRECISION",
		"sql.NullInt64":   "BIGINT",
		"sql.NullString":  "VARCHAR",
		"sql.NullBool":    "BOOLEAN",
	},
	"mysql": {
		"float32":         "float",
		"float64":         "float",
		"int":             "int",
		"uint":            "int",
		"int64":           "bigint",
		"uint64":          "bigint",
		"string":          "varchar",
		"time.Time":       "timestamp",
		"bool":            "tinyint",
		"[]byte":          "blob",
		"[]uint8":         "blob",
		"sql.NullFloat32": "float",
		"sql.NullFloat64": "float",
		"sql.NullInt64":   "bigint",
		"sql.NullString":  "varchar",
		"sql.NullBool":    "tinyint",
	},
}

func MatchType(driver string, goType string) (string, int, error) {
	driver = NormalizeDriverName(driver)

	sqlType, ok := TypeDict[driver][goType]

	if !ok {
		return "", -1, errors.New(fmt.Sprintf("[crud] Can't match Go type '%s' (%s) with any SQL type.", goType, driver))
	}

	if len, ok := TypeLength[driver][sqlType]; ok {
		return sqlType, len, nil
	}

	return sqlType, -1, nil

}

func NormalizeDriverName(driver string) string {
	switch driver {
	case "postgresql", "pgx", "postgres":
		return "postgres"
	default:
		return "mysql"
	}
}
