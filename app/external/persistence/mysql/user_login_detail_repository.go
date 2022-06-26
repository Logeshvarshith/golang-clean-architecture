package mysql

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user/out"
)

type UserLoginDetailRepository struct {
	db *gorm.DB
}

func NewUserLoginDetailRepository(db *gorm.DB) *UserLoginDetailRepository {
	return &UserLoginDetailRepository{
		db: db,
	}
}

func (r *UserLoginDetailRepository) GetUserLoginDetail(ctx context.Context, filter in.UserLoginDetailEmployeeIDFilter) (userLoginDetail out.UserLoginDetail, err error) {

	tx := r.db.WithContext(ctx)

	fmt.Printf("%v\n", filter)

	db := tx.Table("user_login_details").Select("employee_id", "domain_name", "email_id", "password", "uuid", "isSignedup", "role").Where("employee_id=?", filter.EmployeeID).Scan(&userLoginDetail)

	err = db.Error

	return

}
