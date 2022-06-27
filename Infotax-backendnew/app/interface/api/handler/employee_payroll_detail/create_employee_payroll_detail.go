package employee_payroll_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
)

type EmployeePayrollDetailHandler struct {
	employeePayrollDetailUseCase employee_payroll_detail.UseCaser
}

// NewUserLoginDetailHandler function is used to make new UserLoginDetailHandler struct.
func NewEmployeePayrollDetailHandler(employeePayrollDetailUseCase employee_payroll_detail.UseCaser) EmployeePayrollDetailHandler {
	return EmployeePayrollDetailHandler{
		employeePayrollDetailUseCase: employeePayrollDetailUseCase,
	}
}

func (u *EmployeePayrollDetailHandler) CreateEmployeePayrollDetail(ctx *gin.Context) {

	var detail in.CreateEmployeePayroll
	if ok := handler.ValidateData(ctx, &detail); !ok {
		return
	}

	// Defaul values
	detail.TaxRegime = "0"

	saveRes, err := u.employeePayrollDetailUseCase.CreateEmployeePayrollDetail(ctx, detail)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, saveRes)

}
