package user_login_detail

import (
	"github.com/gin-gonic/gin"
)

// DownloadUsersLoginDetail godoc
// @Summary Return user login details in csv file format
// @Description Return user login details in csv file format
// @Tags User Login Detail
// @ID download_users_login_detail
// @Produce application/octet-stream
// @Success 200 {file} file
// @Failure 404 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /infotax/user_login_detail/download [get]
func (u *UserLoginDetailHandler) DownloadUsersLoginDetail(ctx *gin.Context) {

	fileName, filePath, err := u.userLoginDetailUseCase.DownloadUsersLoginDetail(ctx)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "inline;filename="+fileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")
	ctx.File(filePath)
}
