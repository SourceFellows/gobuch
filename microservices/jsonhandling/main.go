package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	ID        int    `json:"id,omitempty"`
	Firstname string `json:"first,omitempty"`
	Lastname  string `json:"last,omitempty"`
}

func main() {
	customer := &Customer{ID: 1,
		Firstname: "Hans", Lastname: "Wurst"}
	bites, _ := json.Marshal(customer)
	fmt.Println(string(bites))

	customer2 := &Customer{}
	json.Unmarshal(bites, customer2)
	fmt.Println(customer2.Firstname)
}
