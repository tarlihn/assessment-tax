// database/connection.go

package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB
}

// ConnectDB establishes a connection to the PostgreSQL database using the DATABASE_URL environment variable.
func ConnectDB() (*Postgres, error) {
	// Retrieve DATABASE_URL from environment
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable" //for test
		// return fmt.Errorf("DATABASE_URL environment variable not set")
	}

	// Open a connection to the database
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging the database: %v", err)
	}

	return &Postgres{Db: db}, nil
}
