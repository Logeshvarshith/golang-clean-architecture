package employee_payroll_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

// GetAllUserLoginDetail method is used to fetch all the user login details.
func (uc *useCase) GetEmployeePayrollDetail(ctx context.Context) (employeePayrollDeatils []out.EmployeePayrollDetail, err *error.Error) {

	employeePayrollDeatils, rerr := uc.employeePayrollDetailRepo.GetEmployeePayrollDetail(ctx)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	return
}
