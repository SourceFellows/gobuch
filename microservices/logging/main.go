package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	// zusätzliche Felder an log.Entry anhängen
	log.WithFields(log.Fields{
		"importId": "ka18s",
		"size":     10,
	}).Info("Starte den Import")
	// Wiederverwenden von log.Entry damit alle weiteren
	// Aufrufe ebenfalls Felder bekommen
	contextLogger := log.WithFields(log.Fields{
		"importId": "0815s",
		"other":    "Ich werde geloggt",
	})
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
