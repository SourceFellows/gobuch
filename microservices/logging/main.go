package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Testing struct {
	Val1 bool
	Val2 string
	Val3 time.Time
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})

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
		"test":     Testing{Val3: time.Now()},
	})
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
