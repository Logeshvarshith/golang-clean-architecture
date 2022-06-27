package employee_payroll_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
)

func (u *EmployeePayrollDetailHandler) DeleteEmployeePayrollDetail(ctx *gin.Context) {

	var filter in.DeleteEmployeePayroll
	if ok := handler.ValidateData(ctx, &filter); !ok {
		return
	}

	delRes, err := u.employeePayrollDetailUseCase.DeleteEmployeePayrollDetail(ctx, filter)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, delRes)

}
