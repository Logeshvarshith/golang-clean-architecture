package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

type EmployeePayrollDetailDetailRepository interface {
	CreateEmployeePayrollDetail(ctx context.Context, detail in.CreateEmployeePayroll) error
	GetEmployeePayrollDetail(ctx context.Context) ([]out.EmployeePayrollDetail, error)
	DeleteEmployeePayrollDetail(ctx context.Context, filter in.DeleteEmployeePayroll) (int64, error)
}
