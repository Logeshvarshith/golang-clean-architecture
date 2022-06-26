package registry

import (
	"context"

	"github.com/google/wire"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/student"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user"
)

func InjectedUserLoginDetailUseCase1(ctx context.Context) user.UseCaser {
	wire.Build(UserLoginDetailUseCaseSet)
	return nil
}

func InjectedStudentDetailUseCase1(ctx context.Context) student.UseCaser {
	wire.Build(StudentDetailUseCaseSet)
	return nil
}
