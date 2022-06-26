package handler

import (
	"github.com/gin-gonic/gin"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/student"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/student/in"
)

type StudentDetailHandler struct {
	studentUseCase student.UseCaser
}

func NewStudentDetailHandler(studentUseCase student.UseCaser) StudentDetailHandler {
	return StudentDetailHandler{
		studentUseCase: studentUseCase,
	}
}

func (u *StudentDetailHandler) AddStudentDetail(ctx *gin.Context) {
	var filter in.AddStudentDetail
	if ok := BindData(ctx, &filter); !ok {
		return
	}

	u.studentUseCase.AddStudentDetail(ctx, filter)

}
