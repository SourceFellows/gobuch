package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	p := &Person{}
	p.Name = "Kristian"
	(*p).Name = "Kristian" //beide Zeilen sind gleichbedeutend
	fmt.Println(p)
}
