package out

type EmployeePayrollDetail struct {
	EmployeeID        string `json:"employee_id"`
	PanNumber         string `json:"pan_number"`
	UanNumber         string `json:"uan_number"`
	BankAccountNumber string `json:"bank_account_number" `
	BankIfscCode      string `json:"bank_ifsc_code" `
	PassportNumber    string `json:"passport_number" `
	PfAccountNumber   string `json:"pf_account_number" `
	TaxRegime         string `json:"tax_regime" `
	EffectiveFrom     string `json:"effective_from" `
	EpsAccountNumber  string `json:"eps_account_number" `
	PrAccountNumber   string `json:"pr_account_number" `
	EsiNumber         string `json:"esi_number"`
}

type DeleteEmployeePayroll struct {
	EmployeeId        string `json:"employee_id"`
	UanNumber         string `json:"uan_number"`
	BankAccountNumber string `json:"bank_account_number" `
	BankIfscCode      string `json:"bank_ifsc_code"`
	PassportNumber    string `json:"passport_number"`
	PfAccountNumber   string `json:"pf_account_number"`
	TaxRegime         string `json:"tax_regime"`
	EffectiveFrom     string `json:"effective_from"`
	EpsAccountNumber  string `json:"eps_account_number"`
	PrAccountNumber   string `json:"pr_account_number"`
	EsiNumber         string `json:"esi_number"`
}

type SaveResponse struct {
	IsSaved string `json:"isSaved"`
}

type DeleteResponse struct {
	IsDeleted string `json:"isDeleted"`
}
