package main

import (
	"fmt"

	"golang.source-fellows.com/samples/unittest/math"
)

func main() {
	res := math.Add(1, 2)
	fmt.Printf("Das Ergbnis ist %v", res)
}
