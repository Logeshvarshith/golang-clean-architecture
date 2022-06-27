package user_login_detail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"www.ivtlinfoview.com/infotax/infotax-backend/app/interface/api/handler"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/out"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/error"
	"www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/in"
	uldMock "www.ivtlinfoview.com/infotax/infotax-backend/app/usecase/user_login_detail/mock"
)

func TestUserLoginDetailHandler_GetUserLoginDetail(t *testing.T) {
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

	invalidPayload := in.EmployeeIDFilter{
		EmployeeID: "",
	}

	mockUserLoginDetailResp := out.UserLoginDetail{
		DomainName: "saravase",
		EmployeeID: "2521",
		EmailID:    "saravase@gmail.com",
		IsSignedUp: 1,
		Password:   "saravase",
		Role:       "ADMIN",
		UUID:       "1234",
	}

	t.Run("Verify invalid content type ", func(t *testing.T) {
		// create usecase using mock
		useCaser := uldMock.NewMockUseCaser(mockCtrl)
		// mock a method in usecase file
		// useCaser.EXPECT().GetUserLoginDetail(gomock.AssignableToTypeOf(&gin.Context{}), invalidFilter).Return(out.UserLoginDetail{}, fmt.Errorf("Employee ID not found"))
		// cteate a h using mocked usecase
		h := NewUserLoginDetailHandler(useCaser)

		// form the api request
		router := gin.Default()
		router.POST("/test_user_login_detail_api", h.GetUserLoginDetail)

		reqBody, err := json.Marshal(invalidFilter)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/test_user_login_detail_api", bytes.NewBuffer(reqBody))
		// call api request
		router.ServeHTTP(w, r)
		// get the api result
		result := w.Result()
		defer result.Body.Close()

		if assert.Equal(t, http.StatusUnsupportedMediaType, result.StatusCode) {

			respBody, err := json.Marshal(gin.H{
				"error": error.NewUnsupportMediaType(fmt.Sprint("/test_user_login_detail_api only accepts Content-Type application/json")),
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}

	})

	t.Run("Validate request payload ", func(t *testing.T) {
		// create usecase using mock
		useCaser := uldMock.NewMockUseCaser(mockCtrl)
		// mock a method in usecase file
		// useCaser.EXPECT().GetUserLoginDetail(gomock.AssignableToTypeOf(&gin.Context{}), invalidFilter).Return(out.UserLoginDetail{}, fmt.Errorf("Employee ID not found"))
		// cteate a h using mocked usecase
		h := NewUserLoginDetailHandler(useCaser)

		// form the api request
		router := gin.Default()
		router.POST("/test_user_login_detail_api", h.GetUserLoginDetail)

		reqBody, err := json.Marshal(invalidPayload)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/test_user_login_detail_api", bytes.NewBuffer(reqBody))
		r.Header.Set("Content-Type", "application/json")
		// call api request
		router.ServeHTTP(w, r)
		// get the api result
		result := w.Result()
		defer result.Body.Close()

		if assert.Equal(t, http.StatusBadRequest, result.StatusCode) {
			e := error.NewBadRequest("Invalid request parameters. Verify request parameters once.")
			respBody, err := json.Marshal(gin.H{
				"error": e,
				"invalidArgs": []handler.InvalidArgument{
					handler.InvalidArgument{
						Field: "EmployeeID",
						Tag:   "required",
					},
				},
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}

	})

	t.Run("Fetching user login detail with invalid employee ID", func(t *testing.T) {
		// create usecase using mock
		useCaser := uldMock.NewMockUseCaser(mockCtrl)
		// mock a method in usecase file
		useCaser.EXPECT().GetUserLoginDetail(gomock.AssignableToTypeOf(&gin.Context{}), invalidFilter).Return(out.UserLoginDetail{}, error.NewNotFound("employee_id", invalidFilter.EmployeeID))
		// cteate a h using mocked usecase
		h := NewUserLoginDetailHandler(useCaser)

		// form the api request
		router := gin.Default()
		router.POST("/test_user_login_detail_api", h.GetUserLoginDetail)

		reqBody, err := json.Marshal(invalidFilter)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/test_user_login_detail_api", bytes.NewBuffer(reqBody))
		r.Header.Set("Content-Type", "application/json")
		// call api request
		router.ServeHTTP(w, r)
		// get the api result
		result := w.Result()
		defer result.Body.Close()

		if assert.Equal(t, http.StatusNotFound, result.StatusCode) {

			respBody, err := json.Marshal(gin.H{
				"error": error.NewNotFound("employee_id", invalidFilter.EmployeeID),
			})
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}

	})

	t.Run("Fetching user login detail with valid employee ID", func(t *testing.T) {
		// create usecase using mock
		useCaser := uldMock.NewMockUseCaser(mockCtrl)
		// mock a method in usecase file
		useCaser.EXPECT().GetUserLoginDetail(gomock.AssignableToTypeOf(&gin.Context{}), validFilter).Return(mockUserLoginDetailResp, nil)
		// cteate a h using mocked usecase
		h := NewUserLoginDetailHandler(useCaser)

		// form the api request
		router := gin.Default()
		router.POST("/test_user_login_detail_api", h.GetUserLoginDetail)

		reqBody, err := json.Marshal(validFilter)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("POST", "/test_user_login_detail_api", bytes.NewBuffer(reqBody))
		r.Header.Set("Content-Type", "application/json")
		// call api request
		router.ServeHTTP(w, r)
		// get the api result
		result := w.Result()
		defer result.Body.Close()

		if assert.Equal(t, http.StatusOK, result.StatusCode) {

			respBody, err := json.Marshal(mockUserLoginDetailResp)
			assert.NoError(t, err)
			assert.Equal(t, respBody, w.Body.Bytes())
		}

	})

}
