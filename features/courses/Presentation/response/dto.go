package response

import (
	"RestfulAPIElearningVideo/features/courses"
	"database/sql"
	"time"
)

type CreateCourse struct {
	ID        uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	Title string `json:"title"`
	Description string `json:"description"`
	PlaylistID string `json:"playlist_id"`
	//Video []CreateVideo `json:"video"`
}

type DetailCreateResponse struct {
	PlaylistID string `json:"playlist_id"`
}

func ToCourseResponse(c courses.CourseCore) CreateCourse {
	return CreateCourse{
		PlaylistID: c.PlaylistID,
	}
}