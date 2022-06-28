package registry

import (
	"context"

	"github.com/google/wire"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail"
)

func InjectedUserLoginDetailUseCase1(ctx context.Context) user_login_detail.UseCaser {
	wire.Build(UserLoginDetailUseCaseSet)
	return nil
}
