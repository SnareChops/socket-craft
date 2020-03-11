package db

import (
	"reflect"
	"runtime"

	"github.com/jmoiron/sqlx"
)

type MigrationFunc func(*sqlx.Tx) error

type Migration struct {
	Func MigrationFunc
}

func InitMigration(fn MigrationFunc) Migration {
	return Migration{
		Func: fn,
	}
}

func (m *Migration) Run(table string, client Client) {
	name := m.Name()
	println("Running migration " + name + "...")
	tx := client.MustBegin()
	err := m.Func(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	m.Save(table, tx)
	tx.Commit()
}

func (m *Migration) Save(table string, tx *sqlx.Tx) {
	tx.MustExec("INSERT INTO `"+table+"` (`name`,`timestamp`) VALUES (?,CURRENT_TIMESTAMP)", m.Name())
}

func (m *Migration) Name() string {
	return runtime.FuncForPC(reflect.ValueOf(m.Func).Pointer()).Name()
}
