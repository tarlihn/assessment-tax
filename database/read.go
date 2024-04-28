package database

import (
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetPersonalDeduction() (float64, error) {
	// Establish connection to the database
	p, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Db.Close()
	// Database connection successful
	log.Println("Connected to the database successfully!")
	stmt, err := p.Db.Prepare("SELECT personalDeduction FROM allowance WHERE id = $1")
	if err != nil {
		log.Fatal("can't prepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var personalDeduction float64

	err = row.Scan(&personalDeduction)
	if err != nil {
		log.Fatal("can't Scan row into variables1", err)
	}

	fmt.Println("one row", personalDeduction)
	return personalDeduction, nil
}

func GetKReceipt() (float64, error) {
	// Establish connection to the database
	p, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Db.Close()
	// Database connection successful
	log.Println("Connected to the database successfully!")

	var kReceipt float64
	err = p.Db.QueryRow("SELECT kReceipt FROM allowance WHERE id = $1", 1).Scan(&kReceipt)
	if err != nil {
		return 0, fmt.Errorf("failed to query kReceipt maximum: %v", err)
	}
	return kReceipt, nil

}
