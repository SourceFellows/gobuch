package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Customer struct {
	gorm.Model
	FirstName    string
	LastName     string
	CreditCard   CreditCard
	CreditCardId uint
}

type CreditCard struct {
	gorm.Model
	Number string
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
	db.AutoMigrate(&CreditCard{}, &Customer{})

	// Create
	customer := Customer{FirstName: "Hans", LastName: "wurst"}
	customer.CreditCard = CreditCard{Number: "123-123-123"}
	db.Set("gorm:association_autoupdate", false).Set("gorm:association_autoupdate", true).Create(&customer)

	var foundCustomer Customer
	var foundCreditCard CreditCard

	db.Where(&Customer{FirstName: "Hans"}).First(&foundCustomer)
	db.Model(&foundCustomer).Related(&foundCreditCard)

	fmt.Println("Gefunden wurde:", foundCreditCard.Number)

	db.Delete(&foundCustomer)

}
