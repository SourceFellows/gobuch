package main

import (
	"context"
	"log"
	"time"
)

func loop(ctx context.Context) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(c)
				return
			default:
				c <- "Hello World!"
			}
			time.Sleep(1 * time.Second)
		}
	}()
	return c
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for s := range loop(ctx) {
		log.Println(s)
	}

	log.Println("nächste Variante")

	ctx, cancel = context.WithCancel(context.Background())
	time.AfterFunc(5*time.Second, cancel)
	for s := range loop(ctx) {
		log.Println(s)
	}

}
