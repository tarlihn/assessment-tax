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

	// Reduce tax by wht
	tax -= wht
	// var response interface{}
	if tax < 0 {
		return models.RefundResponse{Refund: -tax}, nil
	}

	return models.TextResponse{Tax: tax}, nil

}
