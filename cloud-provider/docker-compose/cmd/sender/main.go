package main

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

type Message struct {
	Text string
}

func sendMessage(url string) error {
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}
	defer nc.Close()
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	return c.Publish("testing", &Message{"Hello World"})
}

func main() {

	url := nats.DefaultURL
	if v, ok := os.LookupEnv("NATS_URL"); ok {
		url = v
	}

	for {
		err := sendMessage(url)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Nachricht versendet")
		}
		time.Sleep(4 * time.Second)
	}

}
