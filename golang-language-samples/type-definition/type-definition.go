package main

import "fmt"

type (
	Point struct{ x, y float64 }
	polar Point
)

func main() {
	var p Point
	fmt.Println(p)
}
