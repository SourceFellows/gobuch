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
	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(2 * time.Hour)

	stmt, err := db.Prepare("select EMPNO, ENAME, HIREDATE from EMP where EMPNO = :p")
	if err != nil {
		log.Fatalf("Fehler bei Query: %v", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(sql.Named("p", 7499))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	log.Println("retri")

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
