package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/SnareChops/socket-craft/server/service"
)

func main() {
	fmt.Printf("%v\n", os.Environ())

	s := &service.StandardService{}

	s.AddSocket(service.Actions{}, nil)
	s.AddDB("__migrations", []service.MigrationFunc{})
	close := s.Start()
	defer close()

	// Wait for CTRL-c or client close while handling events.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	select {
	case <-sigChan:
	case <-s.Done():
		log.Print("Router gone, exiting")
		return
	}
}
