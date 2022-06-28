package user_login_detail

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	repoMock "www.ivtlinfoview.com/infotax/infotax-backend/app/domain/repository/mock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"
)

func TestUseCase_GetUserLoginRole(t *testing.T) {

	// testcase is to run in parallel
	t.Parallel()
	// create mock object
	mockCtrl := gomock.NewController(t)
	// cleanup mock object after all of the test methods were finished
	defer mockCtrl.Finish()

	// Add required variables with values corresponding to test
	validFilter := in.EmployeeIDFilter{
		EmployeeID: "2521",
	}

	invalidFilter := in.EmployeeIDFilter{
		EmployeeID: "252123",
	}

	mockUserLoginRoleResp := out.UserLoginRole{
		Role: "Payroll_User",
	}

	t.Run("DB error", func(t *testing.T) {
		// create usecase using mock
		repo := repoMock.NewMockUserLoginDetailRepository(mockCtrl)

		// mock a method in repository file
		repo.EXPECT().GetUserLoginRole(gomock.AssignableToTypeOf(context.TODO()), invalidFilter).Return(out.UserLoginRole{}, errors.New("DB error"))
		// cteate a handler using mocked usecase
		useCase := NewUseCase(repo)

		ulDtl, err := useCase.GetUserLoginRole(context.TODO(), invalidFilter)
		assert.Error(t, err, error.NewInternal())
		assert.Equal(t, "", ulDtl.Role)

	})

	t.Run("Fetching user login detail with invalid employee ID", func(t *testing.T) {
		// create usecase using mock
		repo := repoMock.NewMockUserLoginDetailRepository(mockCtrl)

		// mock a method in repository file
		repo.EXPECT().GetUserLoginRole(gomock.AssignableToTypeOf(context.TODO()), invalidFilter).Return(out.UserLoginRole{}, nil)
		// cteate a handler using mocked usecase
		useCase := NewUseCase(repo)

		ulDtl, err := useCase.GetUserLoginRole(context.TODO(), invalidFilter)
		assert.Error(t, err)
		assert.Equal(t, err, error.NewNotFound("employee_id", invalidFilter.EmployeeID))
		assert.Equal(t, "", ulDtl.Role)

	})

	t.Run("Fetching user login detail with valid employee ID", func(t *testing.T) {
		// create usecase using mock
		repo := repoMock.NewMockUserLoginDetailRepository(mockCtrl)

		// mock a method in repository file
		repo.EXPECT().GetUserLoginRole(gomock.AssignableToTypeOf(context.TODO()), validFilter).Return(mockUserLoginRoleResp, nil)
		// cteate a handler using mocked usecase
		useCase := NewUseCase(repo)

		ulDtl, err := useCase.GetUserLoginRole(context.TODO(), validFilter)
		e := &error.Error{}
		e = nil
		assert.Equal(t, err, e)
		assert.Equal(t, mockUserLoginRoleResp, ulDtl)
	})

}
