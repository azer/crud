package sql

import (
	"errors"
	"fmt"
)

var Types = map[string]int{
	"int":       11,
	"bigint":    20,
	"varchar":   255,
	"text":      -1,
	"date":      -1,
	"time":      -1,
	"timestamp": -1,
}

var TypeDict = map[string]string{
	"int":       "int",
	"string":    "varchar",
	"time.Time": "timestamp",
	"bool":      "tinyint",
}

func MatchType(typeName string) (string, error) {
	if result, ok := TypeDict[typeName]; ok {
		return result, nil
	}

	return "", errors.New(fmt.Sprintf("Unknown type '%s'", typeName))
}
