package employee_official_detail

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/out"
)

// CreateUserLoginDetail method is used to send user login detail to repository and verify the return data from repoistory.
func (uc *useCase) CreateEmployeeOfficialDetail(ctx context.Context, detail in.CreateEmployeeOfficial) (savRes out.SaveResponse, err *error.Error) {

	rerr := uc.employeeOfficialDetailRepo.CreateEmployeeOfficialDetail(ctx, detail)

	serr, ok := rerr.(*mysql.MySQLError)
	if rerr != nil && !ok {
		err = error.NewInternal()
		return
	}

	if ok && serr.Number == 1062 {
		err = error.NewConflict("employee_id", detail.DomainName)
		return
	}

	savRes = out.SaveResponse{
		IsSaved: "true",
	}
	return
}