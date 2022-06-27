package user_login_detail

import (
	"net/http"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
)

// GetUserLoginRole godoc
// @Summary Get login user role based on employee id
// @Description Get role by employee id
// @ID get_user_login_detail_role
// @Accept json
// @Produce json
// @Param message body in.EmployeeIDFilter true "Login user filter"
// @Success 200 {object} out.UserLoginRole
// @Failure 404 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /api/user_login_detail/role [post]
func (u *UserLoginDetailHandler) GetUserLoginRole(ctx *gin.Context) {

	var filter in.EmployeeIDFilter
	if ok := handler.ValidateData(ctx, &filter); !ok {
		return
	}

	role, err := u.userLoginDetailUseCase.GetUserLoginRole(ctx, filter)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, role)

}
