package user_login_detail

import (
	"context"
	"fmt"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// GetUserLoginRole method is used to send filter detail to repository and verify the return data from repoistory.
func (uc *useCase) GetUserLoginRole(ctx context.Context, filter in.EmployeeIDFilter) (userLoginDeatil out.UserLoginRole, err *error.Error) {

	userLoginDeatil, rerr := uc.userLoginDetailRepo.GetUserLoginRole(ctx, filter)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	if userLoginDeatil.Role == "" {
		err = error.NewNotFound("employee_id", filter.EmployeeID)
		return
	}

	fmt.Println(err)
	return
}
