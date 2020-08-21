package main

import (
	"fmt"
	"log"
	"os"
)

var err error
var f *os.File

func write(text string) {
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(f, text)
}

func main() {
	f, err = os.OpenFile("test.txt", os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	write("Hello world 2")
	write("Hello world 2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
