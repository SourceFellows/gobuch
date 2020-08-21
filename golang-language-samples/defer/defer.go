package main

import "fmt"

func sagA() {
	fmt.Println("A")
}

func sagB() {
	fmt.Println("B")
}

func main() {
	defer sagA()
	sagB()
}
