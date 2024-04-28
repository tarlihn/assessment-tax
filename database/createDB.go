package database

import (
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func createDB() error {
	// Establish connection to the database
	p, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Db.Close()
	// Database connection successful
	log.Println("Connected to the database successfully!")

	createTb := `CREATE TABLE IF NOT EXISTS allowance ( id SERIAL PRIMARY KEY, personalDeduction INT);`
	_, err = p.Db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}
	fmt.Println("create table success")

	row := p.Db.QueryRow("INSERT INTO  allowance (personalDeduction) values ($1)  RETURNING id", 60000)
	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't scan id", err)
		return err
	}
	fmt.Println("insert todo success id : ", id)
	return nil
}
