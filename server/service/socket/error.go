package socket

import (
	"fmt"

	"github.com/gammazero/nexus/v3/wamp"
)

// Error Helper for returning errors from a CALL handler
func Error(err interface{}) Result {
	return Result{
		Err: wamp.URI(fmt.Sprint(err)),
	}
}
