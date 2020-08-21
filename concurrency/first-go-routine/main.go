package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printOut() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Print %v\n", i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func main() {
	go printOut()
	time.Sleep(5 * time.Second)
}
