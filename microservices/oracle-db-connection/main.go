package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/godror/godror"
)

type Angestellter struct {
	id                int
	Name              string
	Einstellungsdatum time.Time
}

func main() {
	db, err := sql.Open("godror", "app/app@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)(HOST=172.17.0.1)(PORT=1521)))(CONNECT_DATA=(SID=ORCLCDB)))")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select EMPNO, ENAME, HIREDATE from EMP where EMPNO = :0", 7499)
	if err != nil {
		log.Fatalf("Fehler bei Query: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		angestellter := &Angestellter{}
		err = rows.Scan(&angestellter.id,
			&angestellter.Name,
			&angestellter.Einstellungsdatum)
		if err != nil {
			log.Fatalf("Fehler beim Laden: %v", err)
		}
		log.Println(angestellter)
	}
}
