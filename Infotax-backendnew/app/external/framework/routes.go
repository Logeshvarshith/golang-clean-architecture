package framework

import (
	"context"
	"time"

	employee_official_detail "www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler/employee_official_detail"
	employee_payroll_detail "www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler/employee_payroll_detail"
	user_login_detail "www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler/user_login_detail"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler/middleware"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/config"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/registry"
)

// Handler is used to initialize all the REST endpoints.
func Handler(conf *config.Config) *gin.Engine {

	r := gin.Default()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.Use(middleware.CORS())

	// Swagger Endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userLoginDetailGroup := r.Group(conf.UserLoginDetailBaseUrl)
	userLoginDetailHandler := user_login_detail.NewUserLoginDetailHandler(
		registry.InjectedUserLoginDetailUseCase(ctx),
	)

	userLoginDetailGroup.POST("/detail", userLoginDetailHandler.GetUserLoginDetail)
	userLoginDetailGroup.POST("/role", userLoginDetailHandler.GetUserLoginRole)
	userLoginDetailGroup.GET("/", userLoginDetailHandler.GetAllUserLoginDetail)
	userLoginDetailGroup.DELETE("/", userLoginDetailHandler.DeleteUserLoginDetail)
	userLoginDetailGroup.POST("/", userLoginDetailHandler.CreateUserLoginDetail)
	// userLoginDetailGroup.PUT("/", userLoginDetailHandler.UpdateUserLoginDetail)

	employeePayrollDetailGroup := r.Group(conf.EmployeePayrollDetailBaseUrl)
	employeePayrollDetailHandler := employee_payroll_detail.NewEmployeePayrollDetailHandler(
		registry.InjectedEmployeePayrollDetailuseCase(ctx),
	)

	employeePayrollDetailGroup.POST("/", employeePayrollDetailHandler.CreateEmployeePayrollDetail)
	employeePayrollDetailGroup.GET("/", employeePayrollDetailHandler.GetEmployeePayrollDetail)
	employeePayrollDetailGroup.DELETE("/", employeePayrollDetailHandler.DeleteEmployeePayrollDetail)

	employeeOfficialDetailGroup := r.Group(conf.EmployeeOfficialDetailBaseUrl)
	employeeOfficialDetailHandler := employee_official_detail.NewEmployeeofficialDetailHandler(
		registry.InjectedEmployeeOfficialDetailUseCase(ctx),
	)

	employeeOfficialDetailGroup.POST("/", employeeOfficialDetailHandler.CreateEmployeeOfficialDetail)

	return r

}
