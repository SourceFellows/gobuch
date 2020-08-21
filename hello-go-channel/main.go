package main

import (
	"fmt"
)

func sayHello(c chan string) {
	fmt.Println("vor")
	c <- "Hello World"
	fmt.Println("nach")
}

func main() {
	c := make(chan string)
	go sayHello(c)
	fmt.Println("vor2")
	text := <-c
	fmt.Println("nach2")
	fmt.Println(text)
}
