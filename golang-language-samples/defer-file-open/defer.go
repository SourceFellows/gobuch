package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Open("testfile")
	if err != nil {
		log.Fatalf("Konnte die Datei nicht öffnen: %v", err)
	}
	defer f.Close()

	f2, err := os.Open("nichtVorhanden")
	if err != nil {
		log.Fatalf("Konnte die Datei nicht öffnen: %v", err)
	}
	defer f2.Close()
	//Arbeit mit den Dateien
}
