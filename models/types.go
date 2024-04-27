package models

type Allowance struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type RequestBody struct {
	TotalIncome float64     `json:"totalIncome"`
	WHT         float64     `json:"wht"`
	Allowances  []Allowance `json:"allowances"`
}

type TextResponse struct {
	Tax float64 `json:"tax"`
}

type RefundResponse struct {
	Refund float64 `json:"refund"`
}
