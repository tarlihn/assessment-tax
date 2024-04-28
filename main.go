package main

import (
	"github.com/tarlihn/assessment-tax/tax"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// // Establish connection to the database
	// db, err := database.ConnectDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// // Database connection successful
	// log.Println("Connected to the database successfully!")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	// tax := e.Group("/tax")
	e.POST("/tax/calculation", tax.TaxController)

	e.Logger.Fatal(e.Start(":1323"))
}
