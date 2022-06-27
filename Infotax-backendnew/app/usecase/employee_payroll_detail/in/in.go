package in

type CreateEmployeePayroll struct {
	EmployeeId        string `json:"employee_id" validate:"required"`
	PanNumber         string `json:"pan_number" validate:"required"`
	UanNumber         string `json:"uan_number" validate:"required"`
	BankAccountNumber string `json:"bank_account_number" validate:"required"`
	BankIfscCode      string `json:"bank_ifsc_code" validate:"required"`
	PassportNumber    string `json:"passport_number" validate:"required"`
	PfAccountNumber   string `json:"pf_account_number" validate:"required"`
	TaxRegime         string `json:"tax_regime" validate:"required"`
	EffectiveFrom     string `json:"effective_from" validate:"required"`
	EpsAccountNumber  string `json:"eps_account_number" validate:"required"`
	PrAccountNumber   string `json:"pr_account_number" validate:"required"`
	EsiNumber         string `json:"esi_number" validate:"required"`
}

type DeleteEmployeePayroll struct {
	EmployeeId string `json:"employee_id" validate:"required"`
}
