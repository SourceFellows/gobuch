package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("datei-ist-nicht-da.txt")
	if err != nil {
		//Durch den Aufruf von log.Fatalf wird die weitere
		// Verarbeitung abgebrochen und die Anwendung beendet
		log.Fatalf("Konnte die Datei nicht öffnen, da %v", err)
	}
	f.Close()
}
