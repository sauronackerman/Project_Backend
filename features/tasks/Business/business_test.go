package Business

import (
	 mocks2 "RestfulAPIElearningVideo/features/courses/mocks"
	"RestfulAPIElearningVideo/features/tasks"
	mocks1 "RestfulAPIElearningVideo/features/tasks/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	tkData mocks1.Data
	tkBusiness tasks.Business
	tkValue tasks.TaskCore
	crsBusiness mocks2.Business
	//tkValue tasks.TaskCore


)

func TestMain(m *testing.M) {
	tkBusiness = NewTaskBusiness(&tkData, &crsBusiness)
	tkValue = tasks.TaskCore{
		ID:          0,
		Description: "",
		VideoId:     "1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	os.Exit(m.Run())

}

func TestCreateTaskByVideoId(t *testing.T)  {
	t.Run("create task sukses", func(t *testing.T) {
		tkData.On("InsertData", mock.Anything).Return(tkValue, nil).Once()
		_, _, status := tkBusiness.CreateTaskByVideoId(tkValue)
		assert.Equal(t, status, http.StatusOK)
	})
	t.Run("create task gagal", func(t *testing.T) {
		tkData.On("InsertData", mock.Anything).Return(tkValue, errors.New("err")).Once()
		_, err, status := tkBusiness.CreateTaskByVideoId(tkValue)
		assert.Equal(t, status, http.StatusInternalServerError)
		assert.Error(t, err)
	})
}

func TestFindTaskByVideoId(t *testing.T) {
	t.Run("find task sukses", func(t *testing.T) {
		tkData.On("SelectTaskByVideoId", mock.AnythingOfType("string")).Return(tkValue, nil).Once()
		data, _, _ := tkBusiness.FindTaskByVideoId(tkValue.VideoId)
		assert.Equal(t, data, tkValue)
	})
	t.Run("find task gagal", func(t *testing.T) {
		tkData.On("SelectTaskByVideoId", mock.AnythingOfType("string")).Return(tkValue, errors.New("err")).Once()
		_, err, status := tkBusiness.FindTaskByVideoId(tkValue.VideoId)
		assert.Equal(t, status, http.StatusInternalServerError)
		assert.Error(t, err)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("delete task sukses", func(t *testing.T) {
		tkData.On("DeleteTask", mock.AnythingOfType("string")).Return(tkValue, nil).Once()
		data, _ := tkBusiness.DeleteTask(tkValue.VideoId)
		assert.Equal(t, data, tkValue)
	})
	t.Run("delete task gagal", func(t *testing.T) {
		tkData.On("DeleteTask", mock.AnythingOfType("string")).Return(tkValue, errors.New("err")).Once()
		_, err := tkBusiness.DeleteTask(tkValue.VideoId)
		assert.Error(t, err)
	})
}