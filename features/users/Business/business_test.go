package Business

import (
	"RestfulAPIElearningVideo/features/users"
	usermock "RestfulAPIElearningVideo/features/users/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"

	//uBusiness "RestfulAPIElearningVideo/features/users/Business"

	"os"
	"time"

	"testing"
)

var (
	userData usermock.Data
	userBusiness users.Business
	//userBusiness users.Business
	userValue users.UserCore
	userCourseValue users.UserCourse
	userVideoValue []users.UserCourseVideo
	userLogin users.UserCore
	usersData []users.UserCore
	riza users.UserCore
	course1 users.UserCourse
	video1 users.UserCourseVideo



)


func TestMain(m *testing.M)  {
	userBusiness = NewUserBusiness(&userData)

	course1 = users.UserCourse{
		ID:        1,
		CourseID:  "123",
		UserID:    1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	riza = users.UserCore{
		ID:          1,
		Name:        "Riza",
		Username:    "riza",
		Password:    "123",
		UserCourses: nil,
		UserTasks:   nil,
		Token:       "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	video1 = users.UserCourseVideo{
		ID:        1,
		UserID:    1,
		CourseID:  "12",
		Title:     "aaa",
		VideoID:   "121",
		Note:      "dw",
		Duration:  "",
		UpdatedAt: time.Time{},
		CreatedAt: time.Time{},
	}

	userValue = users.UserCore{
		ID:          1,
		Name:        "Riza",
		Username:    "riza",
		Password:    "123",
		UserCourses: []users.UserCourse{
			{
				ID:        1,
				CourseID:  "123",
				UserID:    1,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
		},
		UserTasks:   []users.UserTaskCore{
			{
				ID:          1,
				UserID:      1,
				TaskID:      "123",
				UserTaskURL: "aaa",
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
		},
		Token:       "sfse",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	userLogin = users.UserCore{
		Username: "riza",
		Password: "123",
	}
	usersData = []users.UserCore{
		{
			ID:          1,
			Name:        "mahar",
			Username:    "mahar",
			Password:    "123",
			UserCourses: nil,
			UserTasks:   nil,
			Token:       "wew",
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
	}

	userVideoValue = []users.UserCourseVideo{
		{
			ID:        1,
			UserID:    1,
			CourseID:  "123",
			Title:     "e",
			VideoID:   "123",
			Note:      "1",
			Duration:  "1",
			UpdatedAt: time.Time{},
			CreatedAt: time.Time{},
		},
	}
	os.Exit(m.Run())
}

func TestLoginUser(t *testing.T) {
	t.Run("Login sukses", func(t *testing.T) {
		userData.On("CheckUser", mock.AnythingOfType("users.UserCore")).Return(userValue, nil).Once()
		data, err := userBusiness.LoginUser(userLogin)
		assert.Equal(t, userValue.Username, data.Username)
		assert.Nil(t, err)
	})
	t.Run("Login gagal", func(t *testing.T) {
		userData.On("CheckUser", mock.AnythingOfType("users.UserCore")).Return(userValue, nil).Once()
		_, err := userBusiness.LoginUser(userLogin)
		assert.NotEqual(t, userValue.Username, "adad")
		assert.Nil(t, err)
	})


	t.Run("Login invalid", func(t *testing.T) {
		userData.On("CheckUser", mock.Anything).Return(userValue, errors.New("invalid data")).Once()
		data, err := userBusiness.LoginUser(users.UserCore{
			Username: "dawsd",
			Password: "123",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid data")
		assert.Empty(t, data.Username)
	})

	t.Run("Error check user func", func(t *testing.T) {
		userData.On("CheckUser", mock.Anything).Return(users.UserCore{}, errors.New("error check data")).Once()
		data, err := userBusiness.LoginUser(userLogin)
		assert.Equal(t, "error check data", err.Error())
		assert.NotNil(t, err)
		assert.Empty(t, data.ID)
	})
}

func TestUserChooseCourse(t *testing.T)  {
		t.Run("success get course", func(t *testing.T) {
			userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(riza, nil).Once()

			userData.On("SelectCourseByPlaylistId", mock.AnythingOfType("string")).Return("123", nil).Once()

			userData.On("InsertUserCourse", mock.Anything).Return(nil).Once()

			err, status := userBusiness.UserChooseCourse(riza.Username, "123")

			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, status)
		})

		t.Run("failed SelectUserByUsername", func(t *testing.T) {
			userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(users.UserCore{}, errors.New("err")).Once()

			err, status := userBusiness.UserChooseCourse(riza.Username, "123")

			assert.Error(t, err)
			assert.Equal(t, http.StatusInternalServerError, status)
		})

		t.Run("username not found", func(t *testing.T) {
			userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(users.UserCore{}, nil).Once()

			err, status := userBusiness.UserChooseCourse(riza.Username, "123")

			assert.Error(t, err)
			assert.Equal(t, http.StatusNotFound, status)
		})

		t.Run("create usercourse", func(t *testing.T) {
			userData.On("SelectUserByUsername", mock.AnythingOfType("string")).Return(riza, nil).Once()

			userData.On("SelectCourseByPlaylistId", mock.AnythingOfType("string")).Return("123", nil).Once()

			userData.On("InsertUserCourse", mock.Anything).Return(errors.New("err")).Once()

			err, status := userBusiness.UserChooseCourse(riza.Username, "123")

			assert.Error(t, err)
			assert.Equal(t, http.StatusInternalServerError, status)
		})

}

func TestUserStartCourse(t *testing.T) {
	t.Run("success start course ", func(t *testing.T) {
		userData.On("GetPlaylistIdforVideo", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userVideoValue, nil).Once()
		userData.On("SelectUserById", mock.AnythingOfType("uint")).Return(riza, nil).Once()
		data, _, status := userBusiness.UserStartCourse(context.Background(), "123" , riza.ID)
		assert.Equal(t, http.StatusOK, status)
		assert.Equal(t, data, userVideoValue)
	})
	t.Run("failed start course ", func(t *testing.T) {
		userData.On("GetPlaylistIdforVideo", mock.Anything, mock.Anything, mock.AnythingOfType("uint")).Return(userVideoValue, errors.New("err")).Once()
		userData.On("SelectUserById", mock.AnythingOfType("uint")).Return(riza, nil).Once()
		_, err, _ := userBusiness.UserStartCourse(context.Background(), "1234" , riza.ID)
		assert.Error(t, err)
	})
}
//goroutine goroutine stack exceeds 1000000000-byte limit
//func TestUpdateUserNote(t *testing.T) {
//	t.Run("sucess edit usernote", func(t *testing.T) {
//		userData.On("UpdateUserNoteData", mock.AnythingOfType("uint"), mock.AnythingOfType("users.UserCourseVideo")).Return(nil).Once()
//		err := userBusiness.UpdateUserNote(riza.ID,video1)
//		assert.Equal(t, err, nil)
//		//userData.On("UpdateUserNoteData", mock.AnythingOfType("uint"), mock.Anything).Return(nil).Once()
//		//userData.On("UpdateUserNoteData", mock.AnythingOfType("uint"), mock.Anything).Return(nil).Once()
//		//err := userBusiness.UpdateUserNote(1, video1)
//		//assert.Nil(t, err)
//	})
//	t.Run("Update user error", func(t *testing.T) {
//		//userData.On("UpdateUserNoteData", mock.Anything).Return(errors.New("error update user")).Once()
//		//err := userBusiness.UpdateUserNote(1, video1)
//		//assert.NotNil(t, err)
//		//assert.Equal(t, "error update user", err.Error())
//	})
//
//}
