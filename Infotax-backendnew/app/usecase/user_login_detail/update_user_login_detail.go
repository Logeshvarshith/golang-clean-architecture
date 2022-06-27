package user_login_detail

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// CreateUserLoginDetail method is used to send user login detail to repository and verify the return data from repoistory.
func (uc *useCase) UpdateUserLoginDetail(ctx context.Context, detail in.UpdateUserDetail) (updRes out.UpdateResponse, err *error.Error) {

	// rows, rerr := uc.userLoginDetailRepo.UpdateUserLoginDetail(ctx, detail)
	// if rows <= 0 {
	// 	err = error.NewNotFound("employee_id", filter.EmployeeID)
	// 	return
	// }

	// if rerr != nil {
	// 	err = error.NewInternal()
	// 	return
	// }

	// savRes = out.UpdateResponse{
	// 	IsUpdated: "true",
	// }
	return
}
