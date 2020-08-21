package main

import "fmt"

func main() {

	c := make(chan string)

	go func() {
		fmt.Println(<-c)
	}()

	c <- "Hello World!"

	close(c)

}
