package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, "Hello World")
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(f, "Hello World")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
