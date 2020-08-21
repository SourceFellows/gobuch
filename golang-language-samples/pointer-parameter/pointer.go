package main

import "fmt"

type Person struct {
	Name string
}

func sayHello(p interface{}) {
	t, ok := p.(*Person)
	fmt.Println(t)
	fmt.Println(ok)
}

func main() {

	p := &Person{"Kristian"}
	sayHello(p)

}
