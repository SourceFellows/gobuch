package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printOut() <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Print %v", i)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()
	return c
}

func join(in1, in2 <-chan string) <-chan string {
	out := make(chan string)
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
	for i := 1; i < 10; i++ {
		fmt.Println(<-c3)
	}
}
