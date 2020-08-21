package main

import "fmt"

func main() {

	val := 2
	var erg int

	switch {
	case val == 2:
		erg = val * 2
		fallthrough
	case val < 10:
		erg = erg * 2
	case val < 100:
		erg = erg * 3
	default:
		erg = 1
	}

	fmt.Println(erg)

}
