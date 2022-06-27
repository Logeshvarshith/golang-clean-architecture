package user_login_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// DeleteUserLoginDetail method is used to send filter detail to repository and verify the return data from repoistory.
func (uc *useCase) DeleteUserLoginDetail(ctx context.Context, filter in.EmployeeIDFilter) (delRes out.DeleteResponse, err *error.Error) {

	rows, rerr := uc.userLoginDetailRepo.DeleteUserLoginDetail(ctx, filter)
	if rows <= 0 {
		err = error.NewNotFound("employee_id", filter.EmployeeID)
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
