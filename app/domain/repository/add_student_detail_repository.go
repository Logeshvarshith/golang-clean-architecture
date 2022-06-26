package repository

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/student/in"
)

type StudentDetailRepository interface {
	AddStudentDetail(ctx context.Context, filter in.AddStudentDetail)
}
