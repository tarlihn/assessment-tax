package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name        string
		totalIncome float64
		wht         float64
		want        float64
	}{
		{
			name:        "Income less than or equal to 150,000, no withholding tax",
			totalIncome: 150000,
			wht:         0,
			want:        0,
		},
		{
			name:        "Income between 150,001 and 500,000, no withholding tax",
			totalIncome: 300000,
			wht:         0,
			want:        24000,
		},
		{
			name:        "Income between 500,001 and 1,000,000, withholding tax applied",
			totalIncome: 700000,
			wht:         10000,
			want:        55000,
		},
		{
			name:        "Income between 1,000,001 and 2,000,000, withholding tax exceeds tax amount",
			totalIncome: 1500000,
			wht:         200000,
			want:        0,
		},
		{
			name:        "Income greater than 2,000,000, no withholding tax",
			totalIncome: 2500000,
			wht:         0,
			want:        415000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := CalculateTax(test.totalIncome, test.wht)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if got != test.want {
				t.Errorf("Incorrect tax calculation. Got: %v, Want: %f", got, test.want)
			}
		})
	}
}
