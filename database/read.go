package database

import (
	"database/sql"
)

func GetPersonalDeduction(db *sql.DB) (float64, error) {
	var personalDeduction float64
	row := db.QueryRow("SELECT personalDeduction FROM allowance WHERE id = $1", 1)
	err := row.Scan(&personalDeduction)
	if err != nil {
		return 0, err
	}
	return personalDeduction, nil
}
