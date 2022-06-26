package student

import (
	"context"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/student/in"
)

type UseCaser interface {
	AddStudentDetail(ctx context.Context, filter in.AddStudentDetail)
}

type useCase struct {
	studentDetailRepo repository.StudentDetailRepository
}

func NewUseCase(studentDetailRepo repository.StudentDetailRepository) UseCaser {
	return &useCase{
		studentDetailRepo: studentDetailRepo,
	}
}
