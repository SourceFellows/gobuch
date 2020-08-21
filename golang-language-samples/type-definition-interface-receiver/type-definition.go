package main

import "fmt"

type Dog interface {
	Bark()
}

type Dackel struct {
	name string
}

func (d Dackel) Bark() {
	fmt.Printf("Wuff %v\n", d.name)
}

func main() {
	var d Dog
	d = Dackel{"Heino"}
	d.Bark()
}
