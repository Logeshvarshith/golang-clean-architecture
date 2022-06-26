package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/app_error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/in"
)

type UserLoginDetailHandler struct {
	userUseCase user.UseCaser
}

func NewUserLoginDetailHandler(userUseCase user.UseCaser) UserLoginDetailHandler {
	return UserLoginDetailHandler{
		userUseCase: userUseCase,
	}
}

func (u *UserLoginDetailHandler) GetUserLoginDetail(ctx *gin.Context) {
	var filter in.UserLoginDetailEmployeeIDFilter
	if ok := BindData(ctx, &filter); !ok {
		return
	}

	userLoginDtl, err := u.userUseCase.GetUserLoginDetail(ctx, filter)
	if err != nil {
		log.Printf("Could not able to find user : %v. Reason : %v\n", filter.EmployeeID, err)
		e := app_error.NewNotFound("employee_id", filter.EmployeeID)
		ctx.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_login_detail": userLoginDtl,
	})

}
