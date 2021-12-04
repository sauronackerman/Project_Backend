package Data

import (
	"RestfulAPIElearningVideo/features/tasks"
	"time"
)

type Video struct {
	ID       	uint `gorm:"primaryKey"`
	CourseID 	string
	Title 		string
	VideoID 	string `gorm:"not null"`
	Duration 	string
	Task		Task
	UpdatedAt 	time.Time
	CreatedAt 	time.Time
}

type Task struct {
	ID uint	`gorm:"primaryKey"`
	VideoID string
	Description string
	UpdatedAt 	time.Time
	CreatedAt 	time.Time
}


func fromCore(core tasks.TaskCore) Task {

	return Task{
		ID:          core.ID,
		VideoID:     core.VideoId,
		Description: core.Description,
	}
}


func toCore(task Task) tasks.TaskCore {

	return tasks.TaskCore{
		ID:   task.ID,
		VideoId: task.VideoID,
		Description: task.Description,
	}
}