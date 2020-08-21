package main

import (
	"fmt"
	"log"
)

//StatementExecutionError ist ein eigener Fehlertyp
//mit zusätzlicher Information
type StatementExecutionError struct {
	Statement string
	Message   string
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
		}
	}
	return nil
}

func main() {
	err := fuehreStatementAus("v:=b")
	//Prüfung über Simple-Statement mit Type-Assertion
	if e, ok := err.(*StatementExecutionError); ok {
		log.Fatalf("Das Statement '%v' konnte nicht ausgeführt werden.", e.Statement)
	} else {
		fmt.Println("alles ok")
	}
}
