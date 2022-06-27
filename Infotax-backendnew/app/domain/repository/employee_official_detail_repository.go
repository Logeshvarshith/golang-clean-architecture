package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/in"
)

type EmployeeOfficialDetailRepository interface {
	CreateEmployeeOfficialDetail(ctx context.Context, detail in.CreateEmployeeOfficial) error
}
