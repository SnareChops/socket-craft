package socket_test

import (
	"testing"

	"github.com/SnareChops/socket-craft/server/service/socket"
)

func TestInitModel(t *testing.T) {

	result := socket.InitModel(TestActions, TestSubscriptions)

	if len(result.Actions) != 1 {
		t.Error("incorrect number of actions")
	}

	if len(result.Subscriptions) != 1 {
		t.Error("incorrect number of subscriptions")
	}

	for key := range result.Actions {
		if key != "test.action" {
			t.Error("incorrect action name")
		}
	}

	for key := range result.Subscriptions {
		if key != "test.sub" {
			t.Error("incorrect subscription name")
		}
	}
}
