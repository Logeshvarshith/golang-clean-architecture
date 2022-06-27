package employee_payroll_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *EmployeePayrollDetailHandler) GetEmployeePayrollDetail(ctx *gin.Context) {

	employeePayrollDtls, err := u.employeePayrollDetailUseCase.GetEmployeePayrollDetail(ctx)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, employeePayrollDtls)

}
