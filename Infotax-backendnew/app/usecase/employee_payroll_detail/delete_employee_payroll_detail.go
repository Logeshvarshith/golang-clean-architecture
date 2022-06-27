package employee_payroll_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

// DeleteUserLoginDetail method is used to send filter detail to repository and verify the return data from repoistory.
func (uc *useCase) DeleteEmployeePayrollDetail(ctx context.Context, filter in.DeleteEmployeePayroll) (delRes out.DeleteResponse, err *error.Error) {

	rows, rerr := uc.employeePayrollDetailRepo.DeleteEmployeePayrollDetail(ctx, filter)
	if rows <= 0 {
		err = error.NewNotFound("employee_id", filter.EmployeeId)
		return
	}

	if rerr != nil {
		err = error.NewInternal()
		return
	}

	delRes = out.DeleteResponse{
		IsDeleted: "true",
	}
	return
}
