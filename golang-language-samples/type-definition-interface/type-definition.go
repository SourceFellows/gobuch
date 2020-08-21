package main

import "fmt"

type File interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}

func main() {

	var f File
	fmt.Println(f)

}
