package courses

import (
	"context"
	"time"
)

type CourseCore struct {
	ID          uint
	Title       string
	Description string
	PlaylistID  string
	Videos      []VideoCore `gorm:"foreignKey:CourseID;type:longtext;reference:PlaylistID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//PlaylistID = PlaylistID, videorul = videoId

type VideoCore struct {
	ID        uint
	Title     string
	CourseID  string
	VideoID   string `gorm:"not null"`
	Duration  string
	Task      Task
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Task struct {
	ID          uint `gorm:"primaryKey"`
	VideoID     uint
	Description string
}

type Business interface {
	CreateCourse(course CourseCore) (CourseCore, error, int)
	AddVideoToCourse(ctx context.Context, playlistId string) ([]VideoCore, error, int)
	DeleteCourseById(id string) (CourseCore, error)
}

type Data interface {
	InsertCourse(course CourseCore) (CourseCore, error)
	GetPlaylistIdforVideo(ctx context.Context, playlistId string) ([]VideoCore, error)
	SelectCourseById(id uint) error
	SelectVideoByVideoId(videoId string) error
	DeleteCourseDataById(id string) (CourseCore, error)
}
