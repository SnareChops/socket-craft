package socket

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/gammazero/nexus/v3/wamp/crsign"
)

func WampConnectFunc(ctx context.Context, url string, cfg client.Config) (Client, error) {
	return client.ConnectNet(ctx, url, cfg)
}

func open() Client {
	var attempt uint8
	var err error
	var c Client
	for attempt, c, err = connect(attempt); err != nil && attempt < 30; attempt, c, err = connect(attempt) {
		if err != nil {
			log.Print(err)
		}
	}
	if err != nil {
		panic(err)
	}
	return c
}

func connect(attempt uint8) (uint8, Client, error) {
	credentials := GetCredentials()
	cfg := client.Config{
		Realm: credentials.Realm,
		HelloDetails: wamp.Dict{
			"authid": credentials.User,
		},
		AuthHandlers: map[string]client.AuthFunc{
			"wampcra": cra(credentials.Pword),
		},
		Logger: log.New(os.Stdout, "", 0),
	}
	time.Sleep(2 * time.Second)
	println("Attempting to connect to WebSocket router...")
	c, err := client.ConnectNet(context.Background(), credentials.Address, cfg)
	return attempt + 1, c, err
}

func cra(secret string) client.AuthFunc {
	return func(c *wamp.Challenge) (string, wamp.Dict) {
		return crsign.RespondChallenge(secret, c, nil), wamp.Dict{}
	}
}
