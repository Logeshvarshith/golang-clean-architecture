package employee_official_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/in"
)

type EmployeeOfficialDetailHandler struct {
	employeeofficialDetailUseCase employee_official_detail.UseCaser
}

// NewUserLoginDetailHandler function is used to make new UserLoginDetailHandler struct.
func NewEmployeeofficialDetailHandler(employeeofficialDetailUseCase employee_official_detail.UseCaser) EmployeeOfficialDetailHandler {
	return EmployeeOfficialDetailHandler{
		employeeofficialDetailUseCase: employeeofficialDetailUseCase,
	}
}

func (u *EmployeeOfficialDetailHandler) CreateEmployeeOfficialDetail(ctx *gin.Context) {

	var detail in.CreateEmployeeOfficial
	if ok := handler.ValidateData(ctx, &detail); !ok {
		return
	}

	saveRes, err := u.employeeofficialDetailUseCase.CreateEmployeeOfficialDetail(ctx, detail)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, saveRes)

}
