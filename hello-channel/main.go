package main

import "fmt"

func main() {
	c := make(chan string)
	//Wert in Channel schreiben
	c <- "Hello World"
	//Wert aus Channel lesen
	fmt.Println(<-c)
}
