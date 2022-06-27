package in

type CreateEmployeeOfficial struct {
	EmployeeID     int64  `json:"employee_id"`
	ProjectID      int64  `json:"project_id"`
	DomainName     string `json:"domain_name"`
	GradeID        int64  `json:"grade_id"`
	OfficialMailID string `json:"official_mail_id"`
	DateOfJoining  string `json:"date_of_joining"`
	Location       string `json:"location"`
	FloorNumber    string `json:"floor_number"`
	SeatNumber     string `json:"seat_number"`
}
