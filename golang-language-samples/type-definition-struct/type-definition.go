package main

import "fmt"

type Person struct {
	vorname  string
	nachname string
}

func main() {
	p := Person{"Kristian", "Köhler"}
	p2 := Person{nachname: "Köhler"}
	fmt.Println(p.vorname)
	fmt.Println(p2.nachname)
}
