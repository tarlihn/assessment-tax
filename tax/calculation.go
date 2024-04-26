package tax

const personalDeduction = 60000

func CalculateTax(totalIncome float64) (float64, error) {
	// Calculate net income
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

	return tax, nil
}
