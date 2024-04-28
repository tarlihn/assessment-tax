package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	// // Test Case 1: Basic Scenario
	// totalIncome1 := 100000.0
	// wht1 := 5000.0
	// allowances1 := []models.Allowance{} // No allowances

	// expectedRefund1 := 5000.0 // Expected tax amount for the basic scenario

	// result1, err1 := CalculateTax(totalIncome1, wht1, allowances1)
	// assert.NoError(t, err1)
	// assert.IsType(t, models.RefundResponse{}, result1)
	// refundResponse1 := result1.(models.RefundResponse)
	// assert.Equal(t, expectedRefund1, refundResponse1.Refund)

	// // Test Case 2: Scenario with Allowances
	// totalIncome2 := 800000.0
	// wht2 := 30000.0
	// allowances2 := []models.Allowance{
	// 	{AllowanceType: "donation", Amount: 120000},  // Example donation allowance exceeding the maximum
	// 	{AllowanceType: "k-receipt", Amount: 150000}, // Example k-receipt allowance exceeding the maximum
	// }

	// expectedTax2 := 11000.0 // Expected tax amount for the scenario with allowances

	// result2, err2 := CalculateTax(totalIncome2, wht2, allowances2)
	// assert.NoError(t, err2)
	// assert.IsType(t, models.TextResponse{}, result2)
	// taxResponse2 := result2.(models.TextResponse)
	// assert.Equal(t, expectedTax2, taxResponse2.Tax)

	// // Test Case 3:
	// totalIncome3 := 500000.0
	// wht3 := 100000.0
	// allowances3 := []models.Allowance{
	// 	{AllowanceType: "donation", Amount: 200000},
	// }

	// expectedTax3 := 19000.0 //

	// result3, err3 := CalculateTax(totalIncome3, wht3, allowances3)
	// assert.NoError(t, err3)
	// assert.IsType(t, models.TextResponse{}, result3)
	// taxResponse3 := result3.(models.TextResponse)
	// assert.Equal(t, expectedTax3, taxResponse3.Tax)

}
