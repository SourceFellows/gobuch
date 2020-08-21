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
			select {
			case t := <-in1:
				out <- t
			case t := <-in2:
				out <- t
			}
		}
	}()
	return out
}

func main() {
	c1 := printOut()
	c2 := printOut()
	c3 := join(c1, c2)

	timeout := time.After(3 * time.Second)
	for {
		select {
		case t := <-c3:
			fmt.Println(t)
		case <-timeout:
			return
		}
	}
}
