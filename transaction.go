package crud

import (
	stdsql "database/sql"
	"fmt"
)

type Tx struct {
	Client *stdsql.Tx
}

func (tx *Tx) Exec(sql string, params ...interface{}) (stdsql.Result, error) {
	return tx.Client.Exec(sql, params...)
}

func (tx *Tx) Query(sql string, params ...interface{}) (*stdsql.Rows, error) {
	return tx.Client.Query(sql, params...)
}

func (tx *Tx) Commit() error {
	return tx.Client.Commit()
}

func (tx *Tx) Rollback() error {
	return tx.Client.Rollback()
}

func (tx *Tx) Create(record interface{}) error {
	return Create(tx.Exec, record)
}

func (tx *Tx) Read(scanTo interface{}, params ...interface{}) error {
	return Read(tx.Query, scanTo, params)
}

func (tx *Tx) Update(record interface{}) error {
	r, err := Update(tx.Exec, record)
	a, err := r.RowsAffected()
	fmt.Println(a)
	fmt.Println(err)
	return err
}

func (tx *Tx) MustUpdate(record interface{}) error {
	return MustUpdate(tx.Exec, record)
}

func (tx *Tx) Delete(record interface{}) error {
	_, err := Delete(tx.Exec, record)
	return err
}

func (tx *Tx) MustDelete(record interface{}) error {
	return MustDelete(tx.Exec, record)
}
