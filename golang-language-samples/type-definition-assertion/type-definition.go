package main

import "fmt"

func main() {
	//Variable muss ein Interfacetyp sein
	var i interface{} = "hello"

	//direkte Zuweisung ohne Prüfung
	s := i.(string)
	fmt.Println(s)

	//Zuweisung mit Prüfung ob möglich
	s, ok := i.(string)
	fmt.Println(s, ok)

	//Zuweisung mit Prüfung ob möglich
	f, ok := i.(int)
	fmt.Println(f, ok)

	//Zuweisung ohne Prüfung führt zu Panic
	f = i.(int)
	fmt.Println(f)
}
