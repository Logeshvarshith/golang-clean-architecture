package user

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/out"
)

type UseCaser interface {
	GetUserLoginDetail(ctx context.Context, filter in.UserLoginDetailEmployeeIDFilter) (out.UserLoginDetail, error)
}

type useCase struct {
	userLoginDetailRepo repository.UserLoginDetailRepository
}

func NewUseCase(userLoginDetailRepo repository.UserLoginDetailRepository) UseCaser {
	return &useCase{
		userLoginDetailRepo: userLoginDetailRepo,
	}
}
