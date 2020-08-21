package main

import "fmt"

func main() {

	val := 2 * 6

	switch val {
	case 10:
		fmt.Println("10!")
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		fmt.Println("kleiner 10")
	default:
		fmt.Println("größer 10")
	}

}
