package crud

import (
	"errors"
	"fmt"
	"github.com/azer/crud/reflect"
	"github.com/azer/crud/sql"
)

func Read(query QueryFn, scanTo interface{}, allparams []interface{}) error {
	sql, params, err := ResolveReadParams(allparams)
	if err != nil {
		return err
	}

	if !reflect.IsPointer(scanTo) {
		return errors.New("A pointer has to be passed for scanning rows to.")
	}

	if reflect.IsSlice(scanTo) {
		return ReadAll(query, scanTo, sql, params)
	}

	return ReadOne(query, scanTo, sql, params)
}

func ReadOne(query QueryFn, scanTo interface{}, sql string, params []interface{}) error {
	scanner, err := NewScan(scanTo)
	if err != nil {
		return err
	}

	rows, err := query(CompleteSelectQuery(sql, scanner), params...)
	if err != nil {
		return err
	}

	defer rows.Close()

	if err := scanner.One(rows); err != nil {
		return err
	}

	return rows.Err()
}

func ReadAll(query QueryFn, scanTo interface{}, sql string, params []interface{}) error {
	scanner, err := NewScan(scanTo)
	if err != nil {
		return err
	}

	rows, err := query(CompleteSelectQuery(sql, scanner), params...)
	if err != nil {
		return err
	}

	defer rows.Close()

	if err := scanner.All(rows); err != nil {
		return err
	}

	return rows.Err()
}

func ResolveReadParams(params []interface{}) (string, []interface{}, error) {
	if len(params) == 0 {
		return "", []interface{}{}, nil
	}

	var (
		query string
		ok    bool
	)

	if query, ok = params[0].(string); !ok {
		return "", nil, errors.New(fmt.Sprintf("Invalid query: %v", params[0]))
	}

	if len(params) == 1 {
		return query, []interface{}{}, nil
	}

	return query, params[1:], nil
}

func CompleteSelectQuery(query string, scanner *Scan) string {
	if scanner.Table == nil {
		return query
	}

	return sql.CompleteSelectQuery(scanner.Table.SQLName, []string{}, query)
}
