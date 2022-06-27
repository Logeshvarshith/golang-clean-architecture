package user_login_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
)

// DeleteUserLoginDetail godoc
// @Summary Delete login user details based on employee id
// @Description Delete login user details by employee id
// @ID delete_user_login_detail
// @Accept json
// @Produce json
// @Param message body in.EmployeeIDFilter true "Login user filter"
// @Success 200 {object} out.DeleteResponse
// @Failure 404 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /api/user_login_detail [delete]
func (u *UserLoginDetailHandler) DeleteUserLoginDetail(ctx *gin.Context) {

	var filter in.EmployeeIDFilter
	if ok := handler.ValidateData(ctx, &filter); !ok {
		return
	}

	delRes, err := u.userLoginDetailUseCase.DeleteUserLoginDetail(ctx, filter)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, delRes)

}
