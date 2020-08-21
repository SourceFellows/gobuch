package main

import "fmt"

func main() {

	var month = "Januar"
	//Erstellt einen Pointer auf die Variable month
	var monthPtr = &month
	//Ändern den Wert von month über den Pointer monthPtr
	*monthPtr = "Februar"

	fmt.Println(month)

}
