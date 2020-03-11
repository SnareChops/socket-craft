package socket_test

import (
	"testing"

	"github.com/SnareChops/socket-craft/server/service/socket"
)

var TestActions = socket.Actions{
	"test.action": func(invocation *socket.Invocation) socket.Result {
		return socket.Success()
	},
}

var TestSubscriptions = socket.Subscriptions{
	"test.sub": func(event *socket.Event) {},
}

func TestSocketShouldInit(t *testing.T) {
	s := &socket.Socket{}

	close := s.Init(FakeConnectFunc, TestActions, TestSubscriptions)
	if s.Model == nil {
		t.Error("expected instantiated model")
	}
	if s.Client == nil {
		t.Error("expected instantiated client")
	}
	if s.Client.(*TestClient).Closed != false {
		t.Error("should not be prematurely closed")
	}
	close()
	if s.Client.(*TestClient).Closed != true {
		t.Error("should have closed")
	}
	// close := s.Init(TEST_ACTIONS, TEST_SUBSCRIPTIONS)

}
