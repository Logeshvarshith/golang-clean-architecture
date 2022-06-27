package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

type UserLoginDetailRepository interface {
	GetUserLoginDetail(ctx context.Context, filter in.EmployeeIDFilter) (out.UserLoginDetail, error)
	GetUserLoginRole(ctx context.Context, filter in.EmployeeIDFilter) (out.UserLoginRole, error)
	GetAllUserLoginDetail(ctx context.Context) ([]out.UserLoginDetail, error)
	DeleteUserLoginDetail(ctx context.Context, filter in.EmployeeIDFilter) (int64, error)
	CreateUserLoginDetail(ctx context.Context, detail in.CreateUserDetail) error
	UpdateUserLoginDetail(ctx context.Context, detail in.UpdateUserDetail) (int64, error)
}
