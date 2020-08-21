package main

import "fmt"

type QuasselStrippe interface {
	Quassel()
}

type Person struct {
	name string
}

func (p *Person) Quassel() {
	fmt.Printf("Hi. Meine Name ist %s\n", p.name)
}

func VielQuasseln(qs QuasselStrippe) {
	for i := 0; i < 3; i++ {
		qs.Quassel()
	}
}

func main() {
	s := &Person{name: "Quassel-Philip"}
	s.Quassel()
	VielQuasseln(s)
}
