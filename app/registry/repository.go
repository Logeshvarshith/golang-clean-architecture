package registry

import (
	"github.com/google/wire"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/config"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/external/persistence/mysql"
)

var (
	repositorySet = wire.NewSet(
		config.ParseConfig,
		config.NewDB,
		mysql.NewUserLoginDetailRepository,
		wire.Bind(new(repository.UserLoginDetailRepository), new(*mysql.UserLoginDetailRepository)),
	)
)

var (
	studentSet = wire.NewSet(
		config.ParseConfig,
		config.NewDB,
		mysql.NewStudentDetailRepository,
		wire.Bind(new(repository.StudentDetailRepository), new(*mysql.StudentDetailRepository)),
	)
)
