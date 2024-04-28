package tax

import (
	"github.com/tarlihn/assessment-tax/models"
)

const (
	personalDeduction = 60000
	maxDonation       = 100000
	maxKReceipt       = 100000
)

func CalculateTax(totalIncome, wht float64, allowances []models.Allowance) (interface{}, error) {

	taxLevels := []models.TaxLevel{
		{Level: "0-150,000", Tax: 0},
		{Level: "150,001-500,000", Tax: 0},
		{Level: "500,001-1,000,000", Tax: 0},
		{Level: "1,000,001-2,000,000", Tax: 0},
		{Level: "2,000,001 ขึ้นไป", Tax: 0},
	}

	// Calculate net income
	netIncome := totalIncome - personalDeduction

	// Deduct allowances, ensuring they don't exceed the maximum values
	var totalDonation, totalKReceipt float64
	for _, allowance := range allowances {
		switch allowance.AllowanceType {
		case "donation":
			if allowance.Amount > maxDonation {
				allowance.Amount = maxDonation
			}
			totalDonation += allowance.Amount
		case "k-receipt":
			if allowance.Amount > maxKReceipt {
				allowance.Amount = maxKReceipt
			}
			totalKReceipt += allowance.Amount
		}
		netIncome -= allowance.Amount
	}

	// Calculate tax for each bracket
	if netIncome > 0 {
		if netIncome <= 150000 {
			taxLevels[0].Tax = 0
		} else {
			taxLevels[0].Tax = (150000 * 0.0) // 0% for the first bracket
		}

		if netIncome > 150000 {
			taxLevels[1].Tax = (min(netIncome, 500000) - 150000) * 0.1 // 10% for the second bracket
		}

		if netIncome > 500000 {
			taxLevels[2].Tax = (min(netIncome, 1000000) - 500000) * 0.15 // 15% for the third bracket
		}

		if netIncome > 1000000 {
			taxLevels[3].Tax = (min(netIncome, 2000000) - 1000000) * 0.2 // 20% for the fourth bracket
		}

		if netIncome > 2000000 {
			taxLevels[4].Tax = (netIncome - 2000000) * 0.35 // 35% for the fifth bracket
		}
	}

	tax := sumTaxLevels(taxLevels)

	// Reduce tax by wht
	tax -= wht
	// var response interface{}
	if tax < 0 {
		return models.RefundResponse{Refund: -tax, TaxLevel: taxLevels}, nil
	}

	return models.TaxResponse{Tax: tax, TaxLevel: taxLevels}, nil

}

// Function to calculate the sum of tax levels
func sumTaxLevels(levels []models.TaxLevel) float64 {
	totalTax := 0.0
	for _, level := range levels {
		totalTax += level.Tax
	}
	return totalTax
}

// Helper function to get the minimum of two floats
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
