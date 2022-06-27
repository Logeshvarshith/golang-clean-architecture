package mysql

import (
	"context"

	"gorm.io/gorm"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_official_detail/in"
)

type EmployeeOfficialDetailRepository struct {
	db *gorm.DB
}

func NewEmployeeOfficialDetailRepository(db *gorm.DB) *EmployeeOfficialDetailRepository {
	return &EmployeeOfficialDetailRepository{
		db: db,
	}
}

func (r *EmployeeOfficialDetailRepository) CreateEmployeeOfficialDetail(ctx context.Context, detail in.CreateEmployeeOfficial) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_official_mst").Create(&detail)
	err = db.Error
	return
}
