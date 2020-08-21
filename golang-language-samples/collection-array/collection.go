package main

import "fmt"

func main() {
	//Initialisierung eines leeren Arrays
	var str1 [3]string
	//Befüllung durch Zuweisungen
	str1[0] = "Hello"
	str1[1] = "World"
	str1[2] = "!"
	fmt.Println(str1[0])
	fmt.Println(len(str1))

	//Initialsierung mit direkter Befüllung
	var str2 = [3]string{"Hello", "World", "!"}
	fmt.Println(str2[0])
	fmt.Println(len(str2))
}
