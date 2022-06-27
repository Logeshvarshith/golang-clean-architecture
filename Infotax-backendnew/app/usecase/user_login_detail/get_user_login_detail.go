package user_login_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// GetUserLoginDetail method is used to send data to userLoginDetail repository and verify the return data from repoistory.
func (uc *useCase) GetUserLoginDetail(ctx context.Context, filter in.EmployeeIDFilter) (userLoginDeatil out.UserLoginDetail, err *error.Error) {

	userLoginDeatil, rerr := uc.userLoginDetailRepo.GetUserLoginDetail(ctx, filter)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	if userLoginDeatil.Role == "" {
		err = error.NewNotFound("employee_id", filter.EmployeeID)
		return
	}

	return
}
