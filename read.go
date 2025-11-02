package crud

import (
	"errors"
	"fmt"

	"github.com/azer/crud/v3/meta"
	"github.com/azer/crud/v3/pg"
)

func read(driver string, query QueryFn, scanTo interface{}, allparams []interface{}) error {
	sql, params, err := ResolveReadParams(allparams)
	if err != nil {
		return err
	}

	// Convert placeholders for postgres
	if pg.IsPostgres(driver) {
		sql = pg.ConvertQueryPlaceholders(sql)
	}

	if !meta.IsPointer(scanTo) {
		return errors.New("A pointer has to be passed for scanning rows to.")
	}

	if meta.IsSlice(scanTo) {
		return readAll(driver, query, scanTo, sql, params)
	}

	return readOne(driver, query, scanTo, sql, params)
}

func readOne(driver string, query QueryFn, scanTo interface{}, sql string, params []interface{}) error {
	scanner, err := NewScan(driver, scanTo)
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

func readAll(driver string, query QueryFn, scanTo interface{}, sql string, params []interface{}) error {
	scanner, err := NewScan(driver, scanTo)
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
