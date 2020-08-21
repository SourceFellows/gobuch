package main

import (
	"fmt"
	"math/rand"
	"time"
)

var c chan string

func printOut() {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("Print %v", i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func main() {
	c = make(chan string)
	go printOut()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
