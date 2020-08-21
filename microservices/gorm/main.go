package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// LogMode enable
	db.LogMode(true)

	// Migrate the schema
	err = db.AutoMigrate(&Customer{}).Error
	if err != nil {
		log.Fatalf("error migrating ", err)
	}

	// Create
	customer := Customer{FirstName: "Hans", LastName: "wurst"}
	db.Create(&customer)

	var foundCustomer Customer
	// Plain SQL
	db.First(&foundCustomer, "First_Name = ?", "Hans")
	// Template
	db.Where(&Customer{FirstName: "Hans"}).First(&foundCustomer)

	fmt.Println("Gefunden wurde:", foundCustomer.FirstName)

	// Update - update product's price to 2000
	db.Model(&foundCustomer).Update("LastName", "Meiser")

	err = db.Delete(&foundCustomer).Error
	if err != nil {
		panic(err)
	}

}
