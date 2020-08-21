package main

import "fmt"

func main() {

	val := 2 * 5

	switch {
	case val == 10:
		fmt.Println("10!")
	case val < 10:
		fmt.Println("kleiner 10")
	default:
		fmt.Println("größer 10")
	}

}
