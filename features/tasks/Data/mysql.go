package Data

import (
	"RestfulAPIElearningVideo/features/tasks"
	"gorm.io/gorm"
)

type taskData struct {
	Conn *gorm.DB
}

func NewTaskData(conn *gorm.DB) *taskData {
	return &taskData{Conn: conn}
}

func (td *taskData)	InsertData(data tasks.TaskCore) (tasks.TaskCore, error)  {
record := fromCore(data)
err := td.Conn.Create(&record).Error
if err != nil {
	return tasks.TaskCore{}, err
}
return toCore(record), nil
}

func (td *taskData) SelectTaskByVideoId(videoId string) (tasks.TaskCore, error) {
	task := Task{}
	err := td.Conn.Where("video_id = ?", videoId).First(&task).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tasks.TaskCore{}, err
	}
	return toCore(task), nil
}

func (td *taskData) DeleteTask(videoId string) (tasks.TaskCore, error) {
	var task Task
	err := td.Conn.Where("playlistId = ?", videoId).Delete(&task).Error
	if err != nil {
		return tasks.TaskCore{}, err
	}
	return toCore(task), nil
}

