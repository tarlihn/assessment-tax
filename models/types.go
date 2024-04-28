package models

type Allowance struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type UserRequest struct {
	TotalIncome float64     `json:"totalIncome"`
	WHT         float64     `json:"wht"`
	Allowances  []Allowance `json:"allowances"`
}

// TaxResponse struct represents the response body for tax calculations
type TaxResponse struct {
	Tax      float64    `json:"tax"`
	TaxLevel []TaxLevel `json:"taxLevel"`
}

type RefundResponse struct {
	Refund   float64    `json:"refund"`
	TaxLevel []TaxLevel `json:"taxLevel"`
}

// TaxLevel struct represents a tax level with its corresponding tax amount
type TaxLevel struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

type AdminRequest struct {
	Amount float64 `json:"amount"`
}

type AdminResponse struct {
	PersonalDeduction float64 `json:"personalDeduction"`
}

type AdminResponseKReciept struct {
	KReciept float64 `json:"kReciept"`
}

type Error struct {
	Message string `json:"message"`
}
