package socket

import (
	"context"

	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
)

// CallHandler the standard CALL handler function
type CallHandler func(invocation *Invocation) Result

type Action = CallHandler
type Subscription = client.EventHandler
type Actions = map[string]Action
type Subscriptions = map[string]Subscription
type Result = client.InvokeResult
type Invocation = wamp.Invocation
type Event = wamp.Event

type Socket interface {
	StartSocket() func()
	SimpleCall(proc string, args []interface{}, kwargs map[string]interface{}) (*wamp.Result, error)
	Register(proc string, handler CallHandler)
	Subscribe(topic string, handler client.EventHandler)
	Publish(topic string, args wamp.List, kwargs wamp.Dict) error
}

type CallFunc = func(proc string, args []interface{}, kwargs map[string]interface{}) (*wamp.Result, error)
type RegisterFunc = func(proc string, handler CallHandler)
type SubscribeFunc = func(topic string, handler client.EventHandler)
type PublishFunc = func(topic string, args wamp.List, kwargs wamp.Dict) error

// WampSocket a WAMP websocket connection
type WampSocket struct {
	Client
	Actions       Actions
	Subscriptions Subscriptions
}

func (s *WampSocket) Init(actions Actions, subscriptions Subscriptions) func() {
	s.Actions = actions
	s.Subscriptions = subscriptions
	return s.StartSocket()
}

func (s *WampSocket) StartSocket() func() {
	s.Client = open()
	s.registerActions()
	s.registerSubscriptions()
	return func() {
		if err := s.Client.Close(); err != nil {
			panic(err)
		}
	}
}

// SimpleCall Sends a CALL message using the provided parameters
func (s *WampSocket) SimpleCall(proc string, args []interface{}, kwargs map[string]interface{}) (*wamp.Result, error) {
	return s.Call(context.Background(), proc, wamp.Dict{wamp.OptInvoke: wamp.InvokeRoundRobin, wamp.OptTimeout: 2000}, args, kwargs, nil)
}

// Register registers a new wamp endpoint using the provided parameters
func (s *WampSocket) Register(proc string, handler CallHandler) {
	fn := func(ctx context.Context, invocation *wamp.Invocation) client.InvokeResult {
		result := handler(invocation)
		return result
	}
	println("Registering:", proc)
	s.Client.Register(proc, fn, wamp.Dict{wamp.OptDiscloseCaller: true, wamp.OptInvoke: wamp.InvokeRoundRobin})
}

// Subscribe subscribes to a topic using the provided parameters
func (s *WampSocket) Subscribe(topic string, handler client.EventHandler) {
	println("Subscribing to:", topic)
	s.Client.Subscribe(topic, handler, nil)
}

func (s *WampSocket) Publish(topic string, args wamp.List, kwargs wamp.Dict) error {
	return s.Client.Publish(topic, nil, args, kwargs)
}

func (s *WampSocket) registerActions() {
	if s.Actions != nil {
		for proc, handler := range s.Actions {
			s.Register(proc, handler)
		}
	}
}

func (s *WampSocket) registerSubscriptions() {
	if s.Subscriptions != nil {
		for topic, handler := range s.Subscriptions {
			s.Subscribe(topic, handler)
		}
	}
}
