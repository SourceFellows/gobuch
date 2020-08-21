package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Could not connect to server because of %v", err)
	}
	defer nc.Close()
	sub, err := nc.SubscribeSync("foo")
	if err != nil {
		log.Fatalf("Could not subscribe because of %v", err)
	}
	m, err := sub.NextMsg(20 * time.Second)
	if m == nil {
		log.Println("No messages available")
		return
	}
	log.Println(m.Subject)
	log.Println(string(m.Data))
	sub.Unsubscribe()
	sub.Drain()
}
