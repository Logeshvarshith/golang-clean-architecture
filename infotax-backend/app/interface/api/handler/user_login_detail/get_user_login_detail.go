package user_login_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail"
)

type UserLoginDetailHandler struct {
	userLoginDetailUseCase user_login_detail.UseCaser
}

// NewUserLoginDetailHandler function is used to make new UserLoginDetailHandler struct.
func NewUserLoginDetailHandler(userLoginDetailUseCase user_login_detail.UseCaser) UserLoginDetailHandler {
	return UserLoginDetailHandler{
		userLoginDetailUseCase: userLoginDetailUseCase,
	}
}

// GetUserLoginDetail godoc
// @Summary Get login user detail based on employee id
// @Description Get detail by employee id
// @Tags User Login Detail
// @ID get_user_login_detail
// @Accept json
// @Produce json
// @Param emp_id path string true "Employee ID"
// @Success 200 {object} out.UserLoginDetail
// @Failure 400 {object} error.Error
// @Failure 404 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail/{emp_id} [get]
func (u *UserLoginDetailHandler) GetUserLoginDetail(ctx *gin.Context) {
	empID := ctx.Param("emp_id")
	if empID == "" {
		err := error.NewBadRequest("Invalid request query parameter. Verify request query parameter once.")
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	userLoginDtl, err := u.userLoginDetailUseCase.GetUserLoginDetail(ctx, empID)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, userLoginDtl)

}