package database

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tarlihn/assessment-tax/models"
)

func UpdatePersonalDeduction(c echo.Context) error {
	var amount float64
	db := DB

	var reqAdmin models.AdminRequest

	err := c.Bind(reqAdmin)
	if err != nil {
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
	stmt, err := db.Prepare(`UPDATE allowance SET personalDeduction = $1`)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
	}

	if _, err := stmt.Exec(amount); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Error{Message: err.Error()})
	}

	// Return the response with HTTP status OK (200)
	return c.JSON(http.StatusOK, models.AdminResponse{PersonalDeduction: amount})
}
