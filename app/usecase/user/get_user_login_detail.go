package user

import (
	"context"
	"fmt"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/out"
)

func (uc *useCase) GetUserLoginDetail(ctx context.Context, filter in.UserLoginDetailEmployeeIDFilter) (userLoginDeatil out.UserLoginDetail, err error) {
	userLoginDeatil, err = uc.userLoginDetailRepo.GetUserLoginDetail(ctx, filter)
	if err != nil {
		return
	}

	if userLoginDeatil.EmployeeID == "" {
		return userLoginDeatil, fmt.Errorf("Employee id: %s not exist", filter.EmployeeID)
	}

	return userLoginDeatil, nil
}
