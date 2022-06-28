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

// CheckIfUserLoginDetailExists method is used to verify whether given employee_id login detail exists or not.
func (r *UserLoginDetailRepository) CheckIfUserLoginDetailExists(ctx context.Context, empID string) (exist bool, err error) {
	var dtl out.UserLoginDetail
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").First(&dtl, empID)
	err = db.Error
	if dtl.EmailID != "" {
		exist = true
	}
	return
}

// GetUserLoginDetail method is used to fetch record from user_login_details table based on the filter condition.
func (r *UserLoginDetailRepository) GetUserLoginDetail(ctx context.Context, empID string) (userLoginDetail out.UserLoginDetail, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Select("employee_id", "domain_name", "email_id", "password", "uuid", "isSignedup", "role").Where("employee_id=?", empID).Scan(&userLoginDetail)
	err = db.Error
	return
}

// GetUserLoginRole method is used to fetch role from user_login_details table based on the filter condition.
func (r *UserLoginDetailRepository) GetUserLoginRole(ctx context.Context, empID string) (userLoginDetail out.UserLoginRole, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Select("role").Where("employee_id=?", empID).Scan(&userLoginDetail)
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
func (r *UserLoginDetailRepository) DeleteUserLoginDetail(ctx context.Context, empID string) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Where("employee_id=?", empID).Delete(&out.UserLoginDetail{})
	err = db.Error
	return
}

// CreateUserLoginDetail method is used to create user login detail in user_login_details table.
func (r *UserLoginDetailRepository) CreateUserLoginDetail(ctx context.Context, detail in.CreateUserDetail) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Create(&detail)
	err = db.Error
	return
}

// UpdateUserLoginDetail method is used to update user login detail from user_login_details table.
func (r *UserLoginDetailRepository) UpdateUserLoginDetail(ctx context.Context, empID string, detail in.UpdateUserDetail) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Where("employee_id=?", empID).Updates(detail)
	err = db.Error
	return
}

// SearchUserLoginDetail method is used to get user login details from user_login_details table based on filterMap conditions.
func (r *UserLoginDetailRepository) SearchUserLoginDetail(ctx context.Context, filterMap map[string]interface{}) (details []out.UserLoginDetail, err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Where(filterMap).Find(&details)
	err = db.Error
	return
}

// CreateBulkUserLoginDetail method is used to create bulk user login detail in user_login_details table.
func (r *UserLoginDetailRepository) CreateBulkUserLoginDetail(ctx context.Context, details []in.UploadUserDetail) (err error) {
	tx := r.db.WithContext(ctx)
	db := tx.Table("user_login_details").Create(&details)
	err = db.Error
	return
}
