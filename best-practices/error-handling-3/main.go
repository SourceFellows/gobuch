package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func NewWriter(w io.Writer) *writer {
	return &writer{w, nil}
}

type writer struct {
	w   io.Writer
	err error
}

func (w *writer) write(text string) {
	if w.err != nil {
		return
	}
	_, w.err = fmt.Fprintln(w.w, text)
}

func (w *writer) Err() error {
	return w.err
}

func main() {

	f, err := os.OpenFile("test.txt", os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := NewWriter(f)

	w.write("Hello world 3")
	w.write("Hello world 3")

	if err = w.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)

}
