package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Client interface {
	MustExec(string, ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Close() error
	Get(interface{}, string, ...interface{}) error
	Select(interface{}, string, ...interface{}) error
	Exec(string, ...interface{}) (sql.Result, error)
	MustBegin() *sqlx.Tx
}
