package user_login_detail

import (
	"net/http"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
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
// @ID get_user_login_detail
// @Accept json
// @Produce json
// @Param message body in.EmployeeIDFilter true "Login user filter"
// @Success 200 {object} out.UserLoginDetail
// @Failure 404 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /api/user_login_detail/detail [post]
func (u *UserLoginDetailHandler) GetUserLoginDetail(ctx *gin.Context) {

	var filter in.EmployeeIDFilter
	if ok := handler.ValidateData(ctx, &filter); !ok {
		return
	}

	userLoginDtl, err := u.userLoginDetailUseCase.GetUserLoginDetail(ctx, filter)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, userLoginDtl)

}
