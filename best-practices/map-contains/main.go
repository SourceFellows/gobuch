package main

import "fmt"

func main() {

	m := make(map[string]string)
	m["key1"] = "value1"

	if v, ok := m["key2"]; ok {
		fmt.Printf("Wert %v enthalten\n", v)
	}

}
