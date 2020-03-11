package socket_test

import (
	"testing"

	"github.com/SnareChops/socket-craft/server/service/socket"
)

func TestSuccess(t *testing.T) {

	expected := []string{"some", "result"}

	result := socket.Success(expected[0], expected[1])

	if len(result.Args) != 2 {
		t.Error("incorrect args length")
	}

	for i, value := range result.Args {
		if value != expected[i] {
			t.Error("incorrect arg value")
		}
	}
}
