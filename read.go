package crud

import (
	"errors"
	"fmt"

	"github.com/azer/crud/v2/meta"
)

func read(query QueryFn, scanTo interface{}, allparams []interface{}) error {
	sql, params, err := ResolveReadParams(allparams)
	if err != nil {
		return err
	}

	if !meta.IsPointer(scanTo) {
		return errors.New("A pointer has to be passed for scanning rows to.")
	}

	if meta.IsSlice(scanTo) {
		return readAll(query, scanTo, sql, params)
	}

	return readOne(query, scanTo, sql, params)
}

func readOne(query QueryFn, scanTo interface{}, sql string, params []interface{}) error {
	scanner, err := NewScan(scanTo)
	if err != nil {
		return err
	}

	rows, err := query(sql, params...)
	if err != nil {
		return err
	}

	defer rows.Close()

	if err := scanner.One(rows); err != nil {
		return err
	}

	return rows.Err()
}

func readAll(query QueryFn, scanTo interface{}, sql string, params []interface{}) error {
	scanner, err := NewScan(scanTo)
	if err != nil {
		return err
	}

	rows, err := query(sql, params...)
	if err != nil {
		return err
	}

	defer rows.Close()

	if err := scanner.All(rows); err != nil {
		return err
	}

	return rows.Err()
}

// Arguments of the common CRUD functions can start with a query in string type, followed
// by parameters for the query itself. ResolveReadParams takes a list of any type,
// returns a query as a string type, parameters as a slice of interface, and potentially an
// error the last parameter.
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
