package main

import "fmt"

func something(text string, i ...int) int {
	fmt.Println(i[0])
	return i[1]
}

func main() {
	fmt.Println(something("hello", 1, 2))
}
