package db_test

import (
	"os"
	"testing"

	"github.com/SnareChops/socket-craft/server/service/db"
)

const (
	DB_ADDRESS         = "test-db-address"
	DB_NAME            = "test-db-name"
	DB_USER            = "test-db-user"
	DB_PWORD           = "test-db-pword"
	ExpectedConnection = "test-db-user:test-db-pword@tcp(test-db-address)/test-db-name?charset=utf8mb4,utf8&parseTime=True"
)

func TestGetCredentials(t *testing.T) {
	os.Setenv("DB_ADDRESS", DB_ADDRESS)
	os.Setenv("DB_NAME", DB_NAME)
	os.Setenv("DB_USER", DB_USER)
	os.Setenv("DB_PWORD", DB_PWORD)

	result := db.GetCredentials()

	if result.Address != DB_ADDRESS {
		t.Error("incorrect db address")
	}

	if result.Name != DB_NAME {
		t.Error("incorrect db name")
	}

	if result.User != DB_USER {
		t.Error("incorrect db user")
	}

	if result.Pword != DB_PWORD {
		t.Error("incorrect db pword")
	}

	if result.Connection != ExpectedConnection {
		t.Error("incorrect connection string")
	}
}
