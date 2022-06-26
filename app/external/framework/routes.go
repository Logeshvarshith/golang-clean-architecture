package framework

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/config"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/registry"
)

func Handler(conf *config.Config) *gin.Engine {

	r := gin.Default()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userLoginDetailGroup := r.Group(conf.UserLoginDetailBaseUrl)
	userLoginDetailHandler := handler.NewUserLoginDetailHandler(
		registry.InjectedUserLoginDetailUseCase(ctx),
	)

	userLoginDetailGroup.POST("/", userLoginDetailHandler.GetUserLoginDetail)

	studentDetailGroup := r.Group(conf.StudentDetailBaseUrl)
	studentDetailHandler := handler.NewStudentDetailHandler(
		registry.InjectedStudentDetailUseCase(ctx),
	)

	studentDetailGroup.POST("/", studentDetailHandler.AddStudentDetail)

	return r

}
