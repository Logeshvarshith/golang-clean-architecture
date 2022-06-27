package mysql

import (
	"context"

	"gorm.io/gorm"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

type UserLoginDetailRepository struct {
	db *gorm.DB
}

// NewUserLoginDetailRepository function is used to make UserLoginDetailRepository struct.
func NewUserLoginDetailRepository(db *gorm.DB) *UserLoginDetailRepository {
	return &UserLoginDetailRepository{
		db: db,
	}
}

// GetUserLoginDetail method is used to fetch record from user_login_details table based on the filter condition.
func (r *UserLoginDetailRepository) GetUserLoginDetail(ctx context.Context, filter in.EmployeeIDFilter) (userLoginDetail out.UserLoginDetail, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Select("employee_id", "domain_name", "email_id", "password", "uuid", "isSignedup", "role").Where("employee_id=?", filter.EmployeeID).Scan(&userLoginDetail)
	err = db.Error
	return
}

// GetUserLoginRole method is used to fetch role from user_login_details table based on the filter condition.
func (r *UserLoginDetailRepository) GetUserLoginRole(ctx context.Context, filter in.EmployeeIDFilter) (userLoginDetail out.UserLoginRole, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Select("role").Where("employee_id=?", filter.EmployeeID).Scan(&userLoginDetail)
	err = db.Error
	return
}

// GetAllUserLoginDetail method is used to fetch all the user login details from user_login_details table.
func (r *UserLoginDetailRepository) GetAllUserLoginDetail(ctx context.Context) (userLoginDetails []out.UserLoginDetail, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Find(&userLoginDetails)
	err = db.Error
	return

}

// DeleteUserLoginDetail method is used to delete user login detail from user_login_details table based on the filter condition.
func (r *UserLoginDetailRepository) DeleteUserLoginDetail(ctx context.Context, filter in.EmployeeIDFilter) (row int64, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Where("employee_id=?", filter.EmployeeID).Delete(&out.UserLoginDetail{})
	err = db.Error
	row = db.RowsAffected
	return
}

// CreateUserLoginDetail method is used to create user login detail from user_login_details table.
func (r *UserLoginDetailRepository) CreateUserLoginDetail(ctx context.Context, detail in.CreateUserDetail) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Create(&detail)
	err = db.Error
	return
}

// UpdateUserLoginDetail method is used to update user login detail from user_login_details table.
func (r *UserLoginDetailRepository) UpdateUserLoginDetail(ctx context.Context, detail in.UpdateUserDetail) (row int64, err error) {
	// tx := r.db.WithContext(ctx)
	// db := tx.Table("user_login_details").Where("employee_id=?", filter.EmployeeID)
	// err = db.Error
	return
}
