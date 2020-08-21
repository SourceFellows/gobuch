package main

import "fmt"

func main() {

	text := "Hello"

	//Nur Zuweisung des Index
	for v := range text {
		fmt.Print(v)
	}

	//Zuweisung von Index und Wert
	for i, v := range text {
		fmt.Printf("index %d value %c\n", i, v)
	}

}
