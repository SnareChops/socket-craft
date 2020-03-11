package socket_test

import (
	"errors"
	"testing"

	"github.com/SnareChops/socket-craft/server/service/socket"
)

func TestError(t *testing.T) {

	result := socket.Error("test-error")

	if result.Err != "test-error" {
		t.Error("incorrect error")
	}

	result = socket.Error(1)

	if result.Err != "1" {
		t.Error("should have converted to string")
	}

	result = socket.Error(errors.New("error-type"))

	if result.Err != "error-type" {
		t.Error("should have handled an error object")
	}
}
