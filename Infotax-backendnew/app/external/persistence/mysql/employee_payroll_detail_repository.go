package mysql

import (
	"context"

	"gorm.io/gorm"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/employee_payroll_detail/out"
)

type EmployeePayrollDetailDetailRepository struct {
	db *gorm.DB
}

func NewEmployeePayrollDetailDetailRepository(db *gorm.DB) *EmployeePayrollDetailDetailRepository {
	return &EmployeePayrollDetailDetailRepository{
		db: db,
	}
}

func (r *EmployeePayrollDetailDetailRepository) CreateEmployeePayrollDetail(ctx context.Context, detail in.CreateEmployeePayroll) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").Create(&detail)
	err = db.Error
	return
}

func (r *EmployeePayrollDetailDetailRepository) GetEmployeePayrollDetail(ctx context.Context) (employeePayrollDetails []out.EmployeePayrollDetail, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").Find(&employeePayrollDetails)
	err = db.Error
	return

}

func (r *EmployeePayrollDetailDetailRepository) DeleteEmployeePayrollDetail(ctx context.Context, filter in.DeleteEmployeePayroll) (row int64, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("employee_payroll_mst").Where("employee_id=?", filter.EmployeeId).Delete(&out.DeleteEmployeePayroll{})
	err = db.Error
	row = db.RowsAffected
	return
}
