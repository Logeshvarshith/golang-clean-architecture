package user_login_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// DeleteUserLoginDetail method is used to send filter detail to repository and verify the return data from repoistory.
func (uc *useCase) DeleteUserLoginDetail(ctx context.Context, empID string) (delRes out.DeleteResponse, err *error.Error) {

	exist, rerr := uc.userLoginDetailRepo.CheckIfUserLoginDetailExists(ctx, empID)
	if !exist {
		err = error.NewNotFound("employee_id", empID)
		return
	}
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	rerr = uc.userLoginDetailRepo.DeleteUserLoginDetail(ctx, empID)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	delRes = out.DeleteResponse{
		IsDeleted: "true",
	}
	return
}
