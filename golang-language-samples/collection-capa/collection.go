package main

import "fmt"

func main() {

	//Array mit fester Größe
	var months = [12]string{"Januar",
		"Februar", "März", "April",
		"Mai", "Juni", "Juli", "August",
		"September", "Oktober", "November",
		"Dezember"}

	//Slices mit immer 3 Elementen des Arrays
	q1 := months[:3]
	q2 := months[3:6]
	q3 := months[6:9]
	q4 := months[9:12]

	fmt.Printf("Length %v, Capa %v, Values: %v\n", len(q1), cap(q1), q1)
	fmt.Printf("Length %v, Capa %v, Values: %v\n", len(q2), cap(q2), q2)
	fmt.Printf("Length %v, Capa %v, Values: %v\n", len(q3), cap(q3), q3)
	fmt.Printf("Length %v, Capa %v, Values: %v\n", len(q4), cap(q4), q4)

}
