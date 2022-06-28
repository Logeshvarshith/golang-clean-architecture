package user_login_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// UpdateUserLoginDetail method is used to share user login detail to repository and verify the return data from repoistory.
func (uc *useCase) UpdateUserLoginDetail(ctx context.Context, empID string, detail in.UpdateUserDetail) (updRes out.UpdateResponse, err *error.Error) {

	exist, rerr := uc.userLoginDetailRepo.CheckIfUserLoginDetailExists(ctx, empID)
	if !exist {
		err = error.NewNotFound("employee_id", empID)
		return
	}

	rerr = uc.userLoginDetailRepo.UpdateUserLoginDetail(ctx, empID, detail)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	updRes = out.UpdateResponse{
		IsUpdated: "true",
	}
	return
}
