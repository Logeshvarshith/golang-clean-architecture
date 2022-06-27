package presenter

import (
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/out"
)

type GetUserLoginDetailResponse struct {
	Data out.UserLoginDetail `json:"data"`
} // @name GetUserLoginDetailResponse

type ErrorResponse struct {
	Error error.Error `json:"error"`
} // @name ErrorResponse
