package socket_test

import (
	"context"

	"github.com/SnareChops/socket-craft/server/service/socket"
	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

type TestClient struct {
	socket.Client
	Called map[string]struct {
		Args   wamp.List
		Kwargs wamp.Dict
	}
	Subscribed    []string
	Registrations map[string]client.InvocationHandler
	Closed        bool
}

func (c *TestClient) Call(ctx context.Context, proc string, options wamp.Dict, args wamp.List, kwargs wamp.Dict, progress client.ProgressHandler) (socket.Result, error) {
	if c.Called == nil {
		c.Called = map[string]struct {
			Args   wamp.List
			Kwargs wamp.Dict
		}{}
	}
	c.Called[proc] = struct {
		Args   wamp.List
		Kwargs wamp.Dict
	}{
		Args:   args,
		Kwargs: kwargs,
	}
	return c.Registrations[proc](ctx, &socket.Invocation{Arguments: args, ArgumentsKw: kwargs}), nil
}

func (c *TestClient) Register(proc string, fn client.InvocationHandler, options wamp.Dict) error {
	if c.Registrations == nil {
		c.Registrations = map[string]client.InvocationHandler{}
	}
	c.Registrations[proc] = fn
	return nil
}

func (c *TestClient) Subscribe(topic string, handler client.EventHandler, options wamp.Dict) error {
	c.Subscribed = append(c.Subscribed, topic)
	return nil
}

func (c *TestClient) Close() error {
	c.Closed = true
	return nil
}

func FakeConnectFunc(ctx context.Context, url string, config client.Config) (socket.Client, error) {
	return &TestClient{}, nil
}
