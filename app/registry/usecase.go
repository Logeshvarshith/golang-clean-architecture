package registry

import (
	"github.com/google/wire"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/student"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user"
)

var (
	UserLoginDetailUseCaseSet = wire.NewSet(
		repositorySet,
		user.NewUseCase,
	)
)

var (
	StudentDetailUseCaseSet = wire.NewSet(
		studentSet,
		student.NewUseCase,
	)
)
