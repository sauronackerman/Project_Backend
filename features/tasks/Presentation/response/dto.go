package response

import "RestfulAPIElearningVideo/features/tasks"

type TaskResp struct {
	VideoID string
	Description string
}


func ToTaskResponse(t tasks.TaskCore) TaskResp {
	return TaskResp{
		VideoID:        t.VideoId,
		Description: t.Description,
	}
}