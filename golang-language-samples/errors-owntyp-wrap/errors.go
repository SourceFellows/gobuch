package main

import (
	"errors"
	"fmt"
	"log"
)

var PermissionError = errors.New("keine Berechtigung")

//StatementExecutionError ist ein eigener Fehlertyp
//mit zusätzlicher Information
type StatementExecutionError struct {
	Statement string
	Message   string
	Cause     error
}

func (wse *StatementExecutionError) Unwrap() error {
	return wse.Cause
}

//Implementierung der Error Methode des error Interface
func (wse *StatementExecutionError) Error() string {
	return wse.Message
}

func fuehreStatementAus(statement string) error {
	if statement != "1=1" {
		return &StatementExecutionError{
			Statement: statement,
			Message:   "Statement kann nicht ausgeführt werden",
			Cause:     PermissionError,
		}
	}
	return nil
}

func main() {
	err := fuehreStatementAus("v:=b")
	//Prüfung über errors Package
	if errors.Is(err, PermissionError) {
		log.Fatal("Das Statement konnte nicht ausgeführt werden.")
	} else {
		fmt.Println("alles ok")
	}

	panic("")
}
