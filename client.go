package crud

import stdsql "database/sql"

type Client interface {
	Exec(string, ...interface{}) (stdsql.Result, error)
	Query(string, ...interface{}) (*stdsql.Rows, error)
	Create(interface{}) error
	CreateAndRead(interface{}) error
	Read(interface{}, ...interface{}) error
	Update(interface{}) error
	Delete(interface{}) error
}
