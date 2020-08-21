package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type LoggerFunc func(message string)

func (l LoggerFunc) Log(message string) {
	l(message)
}

func MyLogFunc(message string) {
	fmt.Printf("I log %s\n", message)
}

func myMethodTakesTheLog(l Logger) {
	l.Log("my log message")
}

func main() {
	logger := LoggerFunc(MyLogFunc)
	myMethodTakesTheLog(logger)
}
