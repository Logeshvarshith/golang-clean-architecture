package user_login_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// GetUserLoginRole method is used to send filter detail to repository and verify the return data from repoistory.
func (uc *useCase) GetUserLoginRole(ctx context.Context, empID string) (userLoginDeatil out.UserLoginRole, err *error.Error) {

	exist, rerr := uc.userLoginDetailRepo.CheckIfUserLoginDetailExists(ctx, empID)
	if !exist {
		err = error.NewNotFound("employee_id", empID)
		return
	}
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	userLoginDeatil, rerr = uc.userLoginDetailRepo.GetUserLoginRole(ctx, empID)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	return
}
