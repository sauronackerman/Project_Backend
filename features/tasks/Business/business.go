package Business

import (
	"RestfulAPIElearningVideo/features/courses"
	"RestfulAPIElearningVideo/features/tasks"
	"errors"
	"net/http"
)

type taskBusiness struct {
	taskData tasks.Data
	courseBusiness courses.Business

}

func NewTaskBusiness(tskData tasks.Data, courseBusiness courses.Business) tasks.Business  {
	return &taskBusiness{tskData, courseBusiness}
}

func (tb *taskBusiness) CreateTaskByVideoId(data tasks.TaskCore) (tasks.TaskCore, error, int)  {
	createdTask, err := tb.taskData.InsertData(data)
	if err != nil {
		return data, err, http.StatusInternalServerError
	}
	return createdTask, nil, http.StatusOK
}


func (tb *taskBusiness) FindTaskByVideoId(videoId string) (tasks.TaskCore, error, int) {
	taskData, err := tb.taskData.SelectTaskByVideoId(videoId)
	if err != nil {
		return tasks.TaskCore{}, err, http.StatusInternalServerError
	}
	if taskData.IsNotFound() {
		return taskData, errors.New("Task not found"), http.StatusNotFound
	}

	return taskData, nil, http.StatusOK
}

func (tb *taskBusiness) DeleteTask(videoId string) (tasks.TaskCore, error) {
	task, err := tb.taskData.DeleteTask(videoId)
	if err != nil {
		return tasks.TaskCore{}, err
	}
	return task, nil
}
