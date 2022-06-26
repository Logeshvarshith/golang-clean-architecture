package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/out"
)

type UserLoginDetailRepository interface {
	GetUserLoginDetail(ctx context.Context, filter in.UserLoginDetailEmployeeIDFilter) (out.UserLoginDetail, error)
}
