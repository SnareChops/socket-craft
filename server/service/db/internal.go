package db

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func open(engine string) Client {
	var attempt uint8
	var err error
	var c Client

	for attempt, c, err = connect(attempt, engine); err != nil && attempt < 30; attempt, c, err = connect(attempt, engine) {
		if err != nil {
			log.Print(err)
		}
	}
	if err != nil {
		panic(err)
	}
	return c
}

func connect(attempt uint8, engine string) (uint8, Client, error) {
	credentials := GetCredentials()
	time.Sleep(2 * time.Second)
	println("Attempting DB Connection...")
	c, err := sqlx.Connect(engine, credentials.Connection)
	return attempt + 1, c, err
}
