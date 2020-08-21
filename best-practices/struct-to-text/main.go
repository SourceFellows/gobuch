package main

import "fmt"

type Tester struct {
	Name     string
	Alter    int
	Sonstwas string
}

func main() {

	t := &Tester{Name: "Dingo", Alter: 3, Sonstwas: "Wert"}
	fmt.Printf("%+v\n", t)

}
