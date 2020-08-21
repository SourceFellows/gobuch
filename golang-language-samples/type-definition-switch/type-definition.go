package main

import "fmt"

func checkMyType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("Es ist ein string")
	case int:
		fmt.Println("Es ist ein int")
	default:
		fmt.Println("Es ist weder string noch int")
	}
}

func main() {
	checkMyType("Hello")
	checkMyType(1)
	checkMyType(3.4)
}
