package main

import "fmt"

func main() {

	//Array mit fester Größe
	var months = [12]string{"Januar",
		"Februar", "März", "April",
		"Mai", "Juni", "Juli", "August",
		"September", "Oktober", "November",
		"Dezember"}
	for i, month := range months {
		fmt.Printf("%v: %v\n", i, month)
	}

	q2 := months[3:6]
	for i, month := range q2 {
		fmt.Printf("%v: %v\n", i, month)
	}

}
