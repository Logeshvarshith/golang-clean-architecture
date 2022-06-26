package out

type UserLoginDetail struct {
	EmployeeID string `json:"employee_id"`
	DomainName string `json:"domain_name"`
	EmailID    string `json:"email_id"`
	Password   string `json:"password"`
	UUID       string `json:"uuid"`
	IsSignedUp int    `json:"isSignedup"`
	Role       string `json:"role"`
}
