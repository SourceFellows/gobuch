package main

import (
	"fmt"
	"os"
)

func main() {
	configValue := os.Getenv("HELLO")
	fmt.Printf("Hello %v\n", configValue)
}
