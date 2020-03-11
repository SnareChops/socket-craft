package service

import (
	"github.com/SnareChops/socket-craft/server/service/db"
	"github.com/SnareChops/socket-craft/server/service/healthcheck"
	"github.com/SnareChops/socket-craft/server/service/socket"
)

type Service interface {
	db.DB
	socket.Socket
	Start() func()
	AddSocket(actions Actions, subs Subscriptions)
	AddDB(table string, migrations []db.MigrationFunc)
}

type StandardService struct {
	*db.Mysql
	*socket.WampSocket
}

type Actions = map[string]socket.Action
type Subscriptions = map[string]socket.Subscription
type MigrationFunc = db.MigrationFunc
type ActionHandler = func(*socket.Invocation) socket.Result

// Init Creates a new service using standard socket-craft project settings
func (s *StandardService) Init(actions Actions, subscriptions Subscriptions, migrationTable string, migrations ...db.MigrationFunc) func() {
	s.AddDB(migrationTable, migrations)
	s.AddSocket(actions, subscriptions)
	return s.Start()
}

func (s *StandardService) AddSocket(actions Actions, subs Subscriptions) {
	if s.WampSocket == nil {
		s.WampSocket = &socket.WampSocket{}
	}
	s.WampSocket.Actions = actions
	s.WampSocket.Subscriptions = subs
}

func (s *StandardService) AddDB(table string, migrations []db.MigrationFunc) {
	if s.Mysql == nil {
		s.Mysql = &db.Mysql{}
	}
	s.Mysql.AddMigrations(table, migrations)
}

func (s *StandardService) Start() func() {
	var socketClose func()
	if s.WampSocket != nil {
		socketClose = s.WampSocket.StartSocket()
	}
	var dbClose func()
	if s.Mysql != nil {
		dbClose = s.Mysql.StartDB()
	}
	healthcheck.Init()
	return func() {
		if socketClose != nil {
			socketClose()
		}
		if dbClose != nil {
			dbClose()
		}
	}
}
