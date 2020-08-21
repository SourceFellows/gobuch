package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalf("Could not connect to server because of %v", err)
	}
	defer nc.Close()

	nc.Publish("foo", []byte("Hello World"))

}
