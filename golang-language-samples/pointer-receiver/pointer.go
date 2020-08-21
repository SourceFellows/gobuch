package main

import (
	"fmt"
)

type Person struct {
	alter int
}

func (p *Person) Alter() int {
	return p.alter
}

func (p *Person) GeburtstagFeiern() {
	p.alter = p.alter + 1
}

func main() {
	p := Person{5}
	p.GeburtstagFeiern()
	fmt.Println(p.Alter())
}
