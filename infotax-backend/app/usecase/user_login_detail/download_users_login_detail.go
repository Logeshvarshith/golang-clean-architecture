package user_login_detail

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
)

// DownloadUsersLoginDetail method is used to return users login details in CSV format.
func (uc *useCase) DownloadUsersLoginDetail(ctx context.Context) (fileName, filePath string, err *error.Error) {

	details, err := uc.GetAllUserLoginDetail(ctx)
	if err != nil {
		return
	}

	filePath = "user_login_detail.csv"
	filetmp, ferr := os.Create(filePath)
	defer filetmp.Close()
	fileName = path.Base(filePath)
	if filePath == "" || fileName == "" || ferr != nil {
		err = error.NewNotFound("File", fileName)
		return
	}

	csvWriter := csv.NewWriter(filetmp)
	defer csvWriter.Flush()

	dtlHeader := []string{"domain_name", "email_id", "employee_id", "isSingedup", "password", "role", "uuid"}
	csvWriter.Write(dtlHeader)

	for _, detail := range details {
		dtl := []string{detail.DomainName, detail.EmailID, detail.EmployeeID, fmt.Sprintf("%v", detail.IsSignedUp), detail.Password, detail.Role, detail.UUID}
		csvWriter.Write(dtl)
	}

	return
}
