package tasks

import "time"

type TaskCore struct {
	ID uint
	Description string
	VideoId string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	CreateTaskByVideoId(data TaskCore) (TaskCore, error, int)
	FindTaskByVideoId(videoId string) (TaskCore, error, int)

	//GetTaskById(id int) (TaskCore, error)
	DeleteTask(videoId string) (TaskCore, error)
	//UpdateTask(data TaskCore) error
}

type Data interface {
	InsertData(data TaskCore) (TaskCore, error)
	SelectTaskByVideoId(videoId string) (TaskCore, error)
	DeleteTask(videoId string) (TaskCore, error)
	//GetTaskData(data TaskCore) ([]TaskCore, error)
	//GetTaskDataById(id int) (TaskCore, error)
	//DeleteTaskData(data TaskCore) error
	//UpdateTaskData(data TaskCore) error
}

func (a *TaskCore) IsNotFound() bool {
	return a.ID == 0
}