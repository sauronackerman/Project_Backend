package request

import (

	"RestfulAPIElearningVideo/features/tasks"
	"time"
)

type CreateTask struct {
	VideoID string
	Description string
	UpdatedAt 	time.Time
	CreatedAt 	time.Time
}

func ToCore(req CreateTask) tasks.TaskCore {
	return tasks.TaskCore{
		//ID:          req.ID,
		//CreatedAt:   req.CreatedAt,
		//UpdatedAt:   req.UpdatedAt,
		VideoId:       req.VideoID,
		Description: req.Description,
		//Videos: ToVideoCore(req.Video),
	}
}