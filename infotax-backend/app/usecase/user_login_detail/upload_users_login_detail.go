package user_login_detail

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/go-playground/validator/v10"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

// UploadUsersLoginDetail method is used to share user login detail to repository and verify the return data from repoistory.
func (uc *useCase) UploadUsersLoginDetail(ctx context.Context, file multipart.File) (savRes out.SaveResponse, err *error.Error) {
	details := []in.UploadUserDetail{}
	lineCnt := 0
	validate := validator.New()
	csvReader := csv.NewReader(file)
	for {
		detail, rerr := csvReader.Read()
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			err = error.NewInternal()
			return
		}
		lineCnt++
		if lineCnt != 1 {
			isSignedUp, cerr := strconv.Atoi(detail[3])
			if cerr != nil {
				err = error.NewBadRequest(fmt.Sprintf("Line no: %d, CSV file contain invalid input values", lineCnt))
				return
			}
			dtl := in.UploadUserDetail{
				DomainName:   detail[0],
				EmailID:      detail[1],
				EmployeeID:   detail[2],
				IsSignedUp:   isSignedUp,
				Password:     detail[4],
				Role:         detail[5],
				UUID:         detail[6],
				EnableAccess: "No",
			}

			if verr := validate.Struct(dtl); verr != nil {
				err = error.NewBadRequest(fmt.Sprintf("Line no: %d, Input validation was failed.", lineCnt))
				return
			}
			details = append(details, dtl)
		}
	}

	rerr := uc.userLoginDetailRepo.CreateBulkUserLoginDetail(ctx, details)
	if rerr != nil {
		err = error.NewInternal()
		return
	}

	return
}
