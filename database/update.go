package database

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/tarlihn/assessment-tax/models"
)

func UpdatePersonalDeduction(c echo.Context) error {
	var amount float64
	// Establish connection to the database
	p, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Db.Close()

	reqAdmin := new(models.AdminRequest)

	if err = c.Bind(reqAdmin); err != nil {
		fmt.Println(reqAdmin)
		return c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}
	if reqAdmin.Amount > 100000 {
		amount = 100000
	} else if reqAdmin.Amount < 10000 {
		amount = 10000
	} else {
		amount = reqAdmin.Amount
	}

	// Execute SQL statement to update personalDeduction in the database
	stmt, err := p.Db.Prepare(`UPDATE allowance SET personalDeduction = $1`)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Error{Message: err.Error()})
	}

	if _, err := stmt.Exec(amount); err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Error{Message: err.Error()})
	}

	// Return the response with HTTP status OK (200)
	return c.JSON(http.StatusOK, &models.AdminResponse{PersonalDeduction: amount})
}
