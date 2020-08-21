package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	bites, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", bites)
	fmt.Println(string(bites))

	toWrite := []byte("Hello World!")
	err = ioutil.WriteFile("test2.txt", toWrite, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
