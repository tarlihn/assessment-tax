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
	stmt, err := p.Db.Prepare("SELECT * FROM allowance WHERE id = $1")
	if err != nil {
		log.Fatal("can't prepare query one row statment", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var personalDeduction float64
	var id int

	err = row.Scan(&id, &personalDeduction)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	fmt.Println("one row", id, personalDeduction)
	return personalDeduction, nil
}
