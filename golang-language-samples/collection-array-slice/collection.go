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
	//Liefert "[Januar februar März]"
	fmt.Println(q1)
	fmt.Println(q2)
	fmt.Println(q3)
	fmt.Println(q4)

}
