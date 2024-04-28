package tax

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tarlihn/assessment-tax/database"
)

type Tax struct {
	TotalIncome float64 `json:"totalIncome"`
	Tax         float64 `json:"tax"`
}

func calculateTax(totalIncome float64, wht float64, donation float64) float64 {
	// Calculate net income
	personalDeduction, _ := database.GetPersonalDeduction()
	netIncome := totalIncome - personalDeduction

	// Apply tax rates based on income brackets
	var tax float64
	switch {
	case netIncome <= 150000:
		tax = 0
	case netIncome <= 500000:
		tax = (netIncome - 150000) * 0.10
	case netIncome <= 1000000:
		tax = (350000)*0.10 + (netIncome-500000)*0.15
	case netIncome <= 2000000:
		tax = (350000)*0.10 + (500000)*0.15 + (netIncome-1000000)*0.20
	default:
		tax = (350000)*0.10 + (500000)*0.15 + (1000000)*0.20 + (netIncome-2000000)*0.35
	}

	tax -= wht
	return tax
}

func CalculateTaxFromCSV(c echo.Context) error {
	// Read form file
	fmt.Println("pass0")
	file, err := c.FormFile("taxFile")
	if err != nil {
		return err
	}
	fmt.Println("pass1")
	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	fmt.Println("pass2")
	// Parse CSV file
	reader := csv.NewReader(src)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record
	var taxes []Tax
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Convert CSV record values to appropriate types
		totalIncome, _ := strconv.ParseFloat(record[0], 64)
		wht, _ := strconv.ParseFloat(record[1], 64)
		donation, _ := strconv.ParseFloat(record[2], 64)

		// Calculate tax
		tax := calculateTax(totalIncome, wht, donation)

		// Create Tax struct
		taxEntry := Tax{
			TotalIncome: totalIncome,
			Tax:         tax,
		}
		taxes = append(taxes, taxEntry)
	}

	// Return JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{"taxes": taxes})
}
