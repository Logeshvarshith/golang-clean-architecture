package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

type UserLoginDetailRepository interface {
	GetUserLoginDetail(ctx context.Context, empID string) (out.UserLoginDetail, error)
	GetUserLoginRole(ctx context.Context, empID string) (out.UserLoginRole, error)
	GetAllUserLoginDetail(ctx context.Context) ([]out.UserLoginDetail, error)
	DeleteUserLoginDetail(ctx context.Context, empID string) error
	CreateUserLoginDetail(ctx context.Context, detail in.CreateUserDetail) error
	UpdateUserLoginDetail(ctx context.Context, empID string, detail in.UpdateUserDetail) error
	CheckIfUserLoginDetailExists(ctx context.Context, empID string) (bool, error)
	SearchUserLoginDetail(ctx context.Context, filterMap map[string]interface{}) ([]out.UserLoginDetail, error)
	CreateBulkUserLoginDetail(ctx context.Context, details []in.UploadUserDetail) error
}
