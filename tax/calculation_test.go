package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name        string
		totalIncome float64
		want        float64
	}{
		{
			name:        "Income less than 150000",
			totalIncome: 100000,
			want:        0,
		},
		{
			name:        "Income between 150000 and 500000",
			totalIncome: 300000,
			want:        9000,
		},
		{
			name:        "Income between 500000 and 1000000",
			totalIncome: 700000,
			want:        56000,
		},
		{
			name:        "Income between 1000000 and 2000000",
			totalIncome: 1500000,
			want:        198000,
		},
		{
			name:        "Income above 2000000",
			totalIncome: 2500000,
			want:        464000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := CalculateTax(test.totalIncome)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if got != test.want {
				t.Errorf("Incorrect tax calculation. Got: %f, Want: %f", got, test.want)
			}
		})
	}
}
