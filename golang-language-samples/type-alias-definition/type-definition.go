package main

import "fmt"

type (
	A1 = string // A1 und string bezeichen identische Typen
	A2 = A1     // A2 und A1 bezeichen identische Typen
)

func main() {
	var text A2
	fmt.Println(text)
}
