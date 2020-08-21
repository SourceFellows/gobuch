package main

import (
	"errors"
	"fmt"
)

func erzeugeEinenNeuenFehler(i int) error {
	if i != 42 {
		return fmt.Errorf("Mit %v lösen Sie das Problem nicht", i)
	}
	return nil
}

func erzeugeNochEinenFehler() error {
	return errors.New("Da ist noch ein Problem")
}

func main() {
	err := erzeugeEinenNeuenFehler(13)
	if err != nil {
		fmt.Printf("Es ist ein Fehler aufgetreten: %v\n", err)
	}
	err = erzeugeNochEinenFehler()
	if err != nil {
		fmt.Printf("Es ist ein Fehler aufgetreten: %v\n", err)
	}

}
