package database

import "log"

func InitDB() {
	db, _ := ConnectDB()
	var err error
	createTb := `CREATE TABLE IF NOT EXISTS allowance ( id SERIAL PRIMARY KEY, personalDeduction INT);`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal(err)
	}

	db.QueryRow("INSERT INTO  allowance (personalDeduction) values ($1)  RETURNING id", 60000)
}
