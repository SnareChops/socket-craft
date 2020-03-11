package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	StartDB() func()
	AddMigrations(table string, migrations []MigrationFunc)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(string, ...interface{}) (sql.Result, error)
	MustExec(string, ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	MustBegin() *sqlx.Tx
}

type GetFunc = func(dest interface{}, query string, args ...interface{}) error
type SelectFunc = func(dest interface{}, query string, args ...interface{}) error
type ExecFunc = func(string, ...interface{}) (sql.Result, error)
type MustExecFunc = func(string, ...interface{}) sql.Result
type NamedExecFunc = func(query string, arg interface{}) (sql.Result, error)
type MustBeginFunc = func() *sqlx.Tx

type Mysql struct {
	Client
	*Migrations
}

// Init initalize and return a database connection
func (d *Mysql) Init(table string, migrations ...MigrationFunc) func() {
	return d.StartDB()
}

func (d *Mysql) StartDB() func() {
	d.Client = open("mysql")
	if d.Table != "" {
		d.Migrations.Run(d.Client)
	}
	return func() {
		if err := d.Client.Close(); err != nil {
			panic(err)
		}
	}
}

func (d *Mysql) AddMigrations(table string, migrations []MigrationFunc) {
	d.Migrations = InitMigrations(table, migrations)
}

func (db *Mysql) Get(dest interface{}, query string, args ...interface{}) error {
	err := db.Client.Get(dest, query, args...)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "no rows") {
			dest = nil
			return nil
		}
	}
	return err
}

func (db *Mysql) Select(dest interface{}, query string, args ...interface{}) error {
	err := db.Client.Select(dest, query, args...)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "no rows") {
			dest = nil
			return nil
		}
	}
	return err
}
