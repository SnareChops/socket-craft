package socket

import "github.com/gammazero/nexus/v3/wamp"

// Success Helper for returning args from a CALL handler
func Success(args ...interface{}) Result {
	list := wamp.List{}
	for _, item := range args {
		list = append(list, item)
	}
	return Result{
		Args: list,
	}
}
