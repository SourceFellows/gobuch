package main

import (
	"fmt"
	"math/rand"
	"time"
)

var c chan string

func printOut() {
	for i := 1; i < 10; i++ {
		c <- fmt.Sprintf("Print %v", i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	close(c)
}

func main() {
	c = make(chan string)
	go printOut()

	for {
		t, open := <-c
		if open {
			fmt.Println(t)
		} else {
			break
		}
	}
}
