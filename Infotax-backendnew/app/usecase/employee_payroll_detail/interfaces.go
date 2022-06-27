package employee_payroll_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

type UseCaser interface {
	CreateEmployeePayrollDetail(ctx context.Context, detail in.CreateEmployeePayroll) (out.SaveResponse, *error.Error)
	GetEmployeePayrollDetail(ctx context.Context) ([]out.EmployeePayrollDetail, *error.Error)
	DeleteEmployeePayrollDetail(ctx context.Context, filter in.DeleteEmployeePayroll) (out.DeleteResponse, *error.Error)
}

type useCase struct {
	employeePayrollDetailRepo repository.EmployeePayrollDetailDetailRepository
}

// NewUseCase function is used to make new userCase struct.
func NewUseCase(employeePayrollDetailRepo repository.EmployeePayrollDetailDetailRepository) UseCaser {
	return &useCase{
		employeePayrollDetailRepo: employeePayrollDetailRepo,
	}
}
