package user_login_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// GetAllUserLoginDetail method is used to fetch all the user login details.
func (uc *useCase) GetAllUserLoginDetail(ctx context.Context) (userLoginDeatils []out.UserLoginDetail, err *error.Error) {

	userLoginDeatils, rerr := uc.userLoginDetailRepo.GetAllUserLoginDetail(ctx)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	return
}
