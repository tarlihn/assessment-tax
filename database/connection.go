// database/connection.go

package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDB establishes a connection to the PostgreSQL database using the DATABASE_URL environment variable.
func ConnectDB() error {
	// Retrieve DATABASE_URL from environment
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return fmt.Errorf("DATABASE_URL environment variable not set")
	}

	// Open a connection to the database
	DB, err := sql.Open("postgres", url)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	// Test the connection
	if err := DB.Ping(); err != nil {
		DB.Close()
		return fmt.Errorf("error pinging the database: %v", err)
	}

	return nil
}
