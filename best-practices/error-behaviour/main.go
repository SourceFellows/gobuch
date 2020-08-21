package main

import (
	"log"
	"net"
)

func isTimeout(err error) bool {

	type timeout interface {
		Timeout() bool
	}

	v, ok := err.(timeout)
	if !ok {
		return false
	}

	return v.Timeout()

}

func main() {
	err := callRemoteService()

	if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
		//nochmal versuchen
		log.Printf("it's a unknown network error: %v\n", nerr)
	}

	if err != nil && !isTimeout(err) {
		log.Fatal(err)
	}
}
