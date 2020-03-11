package socket

import (
	"context"

	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

type ConnectFunc = func(context.Context, string, client.Config) (Client, error)

type Client interface {
	Call(context.Context, string, wamp.Dict, wamp.List, wamp.Dict, client.ProgressHandler) (*wamp.Result, error)
	Register(string, client.InvocationHandler, wamp.Dict) error
	Subscribe(string, client.EventHandler, wamp.Dict) error
	Publish(string, wamp.Dict, wamp.List, wamp.Dict) error
	Close() error
	Done() <-chan struct{}
}
