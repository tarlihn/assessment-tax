package tax

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tarlihn/assessment-tax/models"
)

func TaxController(c echo.Context) error {
	req := new(models.UserRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	// Prepare response
	result, _ := CalculateTax(req.TotalIncome, req.WHT, req.Allowances)

	// Prepare response

	return c.JSON(http.StatusOK, result)
}
