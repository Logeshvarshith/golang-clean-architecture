package employee_payroll_detail

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

// CreateUserLoginDetail method is used to send user login detail to repository and verify the return data from repoistory.
func (uc *useCase) CreateEmployeePayrollDetail(ctx context.Context, detail in.CreateEmployeePayroll) (savRes out.SaveResponse, err *error.Error) {

	rerr := uc.employeePayrollDetailRepo.CreateEmployeePayrollDetail(ctx, detail)

	serr, ok := rerr.(*mysql.MySQLError)
	if rerr != nil && !ok {
		err = error.NewInternal()
		return
	}

	if ok && serr.Number == 1062 {
		err = error.NewConflict("employee_id", detail.EmployeeId)
		return
	}

	savRes = out.SaveResponse{
		IsSaved: "true",
	}
	return
}

