package user_login_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

// DeleteUserLoginDetail godoc
// @Summary Delete login user details based on employee id
// @Description Delete login user details by employee id
// @Tags User Login Detail
// @ID delete_user_login_detail
// @Accept json
// @Produce json
// @Param emp_id path string true "Employee ID"
// @Success 200 {object} out.DeleteResponse
// @Failure 400 {object} error.Error
// @Failure 404 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail/{emp_id} [delete]
func (u *UserLoginDetailHandler) DeleteUserLoginDetail(ctx *gin.Context) {

	empID := ctx.Param("emp_id")
	if empID == "" {
		err := error.NewBadRequest("Invalid request query parameter. Verify request query parameter once.")
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	delRes, err := u.userLoginDetailUseCase.DeleteUserLoginDetail(ctx, empID)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, delRes)

}
