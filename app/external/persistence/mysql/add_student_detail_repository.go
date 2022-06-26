package mysql

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/student/in"
)

type StudentDetailRepository struct {
	db *gorm.DB
}

func NewStudentDetailRepository(db *gorm.DB) *StudentDetailRepository {
	return &StudentDetailRepository{
		db: db,
	}
}

func (r *StudentDetailRepository) AddStudentDetail(ctx context.Context, filter in.AddStudentDetail) {

	tx := r.db.WithContext(ctx)

	fmt.Printf("%v\n", filter)

	db := tx.Table("students").Create(&filter)

	_ = db.Error

	return

}
