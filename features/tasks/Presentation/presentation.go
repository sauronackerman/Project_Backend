package Presentation

import (

	"RestfulAPIElearningVideo/features/tasks"
	"RestfulAPIElearningVideo/features/tasks/Presentation/request"
	"RestfulAPIElearningVideo/features/tasks/Presentation/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskPresentation struct {
	taskBusiness tasks.Business
}

func NewTaskPresentation(taskBusiness tasks.Business)	*TaskPresentation  {
	return &TaskPresentation{taskBusiness: taskBusiness}
}
type json map[string]interface{}

func (tp *TaskPresentation) CreateTask(c echo.Context) error {
	var newTask request.CreateTask
	c.Bind(&newTask)
	task, err, status := tp.taskBusiness.CreateTaskByVideoId(request.ToCore(newTask))
	if err != nil {
		return c.JSON(status, json{
			"message": " ",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, json{"courses": task})
}

func (tp *TaskPresentation) GetTaskByVideoId(c echo.Context) error {
	var videoId string
	//req := c.Param("videoId")
	echo.PathParamsBinder(c).String("videoId", &videoId)
	result, err, status := tp.taskBusiness.FindTaskByVideoId(videoId)
	if err != nil {
		return c.JSON(status, json{
			"message": " ",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, json{"task": result})
}


func (tp *TaskPresentation) DeleteTask(c echo.Context) error {
	videoId := c.Param("videoId")

	data, err := tp.taskBusiness.DeleteTask(videoId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data deleted",
		"data":    response.ToTaskResponse(data)})
}