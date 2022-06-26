package student

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/student/in"
)

func (uc *useCase) AddStudentDetail(ctx context.Context, filter in.AddStudentDetail) {
	uc.studentDetailRepo.AddStudentDetail(ctx, filter)

}
