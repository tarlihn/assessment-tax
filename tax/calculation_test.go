package tax

import (
	"testing"

	"github.com/tarlihn/assessment-tax/models"
)

func TestCalculateTax(t *testing.T) {
	// Test case 1: Tax > 0
	totalIncome1 := 500000.0
	wht1 := 25000.0
	response1, err1 := CalculateTax(totalIncome1, wht1)
	if err1 != nil {
		t.Errorf("Test case 1 failed: unexpected error: %v", err1)
	}
	textResponse1, ok1 := response1.(models.TextResponse)
	if !ok1 {
		t.Errorf("Test case 1 failed: expected TextResponse, got %T", response1)
	}
	expectedTax1 := 4000.0 // Expected tax: 500,000 (รายรับ) - 60,0000 (ค่าลดหย่อนส่วนตัว) = 440,000 tax 29,000.00 - 25,000.00 = 4,000
	if textResponse1.Tax != expectedTax1 {
		t.Errorf("Test case 1 failed: expected tax to be %f, got %f", expectedTax1, textResponse1.Tax)
	}

	// Test case 2: Tax == 0
	totalIncome2 := 150000.0
	wht2 := 2000.0
	response2, err2 := CalculateTax(totalIncome2, wht2)
	if err2 != nil {
		t.Errorf("Test case 2 failed: unexpected error: %v", err2)
	}
	refundResponse2, ok2 := response2.(models.RefundResponse)
	if !ok2 {
		t.Errorf("Test case 2 failed: expected RefundResponse, got %T", response2)
	}
	expectedRefund2 := 2000.0 // Expected refund: 150000*0. - 2000
	if refundResponse2.Refund != expectedRefund2 {
		t.Errorf("Test case 2 failed: expected refund to be %f, got %f", expectedRefund2, refundResponse2.Refund)
	}

}
