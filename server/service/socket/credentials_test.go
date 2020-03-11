package socket_test

import (
	"os"
	"testing"

	"github.com/SnareChops/socket-craft/server/service/socket"
)

const (
	NEXUS_ADDRESS = "test-nexus-address"
	NEXUS_REALM   = "test-nexus-realm"
	NEXUS_USER    = "test-nexus-user"
	NEXUS_PWORD   = "test-nexus-pword"
)

func TestGetCredentials(t *testing.T) {
	if err := os.Setenv("NEXUS_ADDRESS", NEXUS_ADDRESS); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("NEXUS_REALM", NEXUS_REALM); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("NEXUS_USER", NEXUS_USER); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("NEXUS_PWORD", NEXUS_PWORD); err != nil {
		t.Error(err)
	}

	result := socket.GetCredentials()

	if result.Address != NEXUS_ADDRESS {
		t.Error("incorrect NEXUS_ADDRESS")
	}

	if result.Realm != NEXUS_REALM {
		t.Error("incorrect NEXUS_REALM")
	}

	if result.User != NEXUS_USER {
		t.Error("incorrect NEXUS_USER")
	}

	if result.Pword != NEXUS_PWORD {
		t.Error("incorrect NEXUS_PWORD")
	}
}
