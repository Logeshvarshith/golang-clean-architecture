package user_login_detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
)

// UpdateUserLoginDetail godoc
// @Summary Update login user details
// @Description Update login user details
// @ID update_user_login_detail
// @Accept json
// @Produce json
// @Param message body in.UpdateUserDetail true "User login details"
// @Success 200 {object} out.UpdateResponse
// @Failure 404 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /api/user_login_detail [put]
func (u *UserLoginDetailHandler) UpdateUserLoginDetail(ctx *gin.Context) {

	var detail in.UpdateUserDetail
	if ok := handler.ValidateData(ctx, &detail); !ok {
		return
	}

	updRes, err := u.userLoginDetailUseCase.UpdateUserLoginDetail(ctx, detail)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, updRes)

}
