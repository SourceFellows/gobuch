package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	text string
	wait chan interface{}
}

func printOut() <-chan Message {
	finish := make(chan interface{})
	c := make(chan Message)
	go func() {
		for i := 1; ; i++ {
			c <- Message{fmt.Sprintf("Print %v", i), finish}
			time.Sleep(time.Duration(rand.Intn(9000)) * time.Millisecond)
			<-finish
		}
	}()
	return c
}

func join(in1, in2 <-chan Message) <-chan Message {
	out := make(chan Message)
	go func() {
		for {
			out <- <-in1
		}
	}()
	go func() {
		for {
			out <- <-in2
		}
	}()
	return out
}

func main() {
	c1 := printOut()
	c2 := printOut()
	c3 := join(c1, c2)
	for i := 1; i < 20; i++ {
		message1 := <-c3
		message2 := <-c3
		fmt.Println(message1.text)
		fmt.Println(message2.text)
		message1.wait <- ""
		message2.wait <- ""
	}
}
