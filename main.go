package main

import (
	"log"

	"github.com/tarlihn/assessment-tax/database" // Import the database package

	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Establish connection to the database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Database connection successful
	log.Println("Connected to the database successfully!")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
