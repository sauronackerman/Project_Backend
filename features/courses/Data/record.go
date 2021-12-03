package Data

import (
	"RestfulAPIElearningVideo/features/courses"
	"time"
)

type Course struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	PlaylistID  string `gorm:"primaryKey"`
	Videos      []Video `gorm:"foreignKey:CourseID;type:longtext;reference:PlaylistID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

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
	VideoID uint
	Description string
}

func toVideoCore(v *Video) courses.VideoCore {
	return courses.VideoCore{
		ID:       v.ID,
		Title:    v.Title,
		VideoID: v.VideoID,
		Duration: v.Duration,
	}
}

func toCourseCore(c *Course) courses.CourseCore {
	//var convertedVideos []courses.VideoCore
	//for _, req := range c.Videos {
	//	convertedVideos = append(convertedVideos, toVideoCore(req))
	//}
	return courses.CourseCore{
		ID:       	 c.ID,
		Title:       c.Title,
		Description: c.Description,
		PlaylistID:   c.PlaylistID,
		Videos:  	 toSliceVideoCore(c.Videos),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func toCore2(c Course) courses.CourseCore {
	return courses.CourseCore{
		ID:       	 c.ID,
		Title:       c.Title,
		Description: c.Description,
		PlaylistID:   c.PlaylistID,
		Videos:  	 toSliceVideoCore(c.Videos),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}



func toSliceVideoCore(v []Video) []courses.VideoCore {
	videos := make([]courses.VideoCore, len(v))

	for i, v := range v {
		videos[i] = toVideoCore(&v)
	}

	return videos
}

//func fromCourseCore(core courses.CourseCore) Course {
//	return Course{
//		ID:          core.ID,
//		Title:       core.Title,
//		Description: core.Description,
//		PlaylistID:   core.PlaylistID,
//		Videos:      fromSliceVideoCore(core.Videos),
//		UpdatedAt:   time.Time{},
//		CreatedAt:   time.Time{},
//	}
//}
//
func fromVideoCore(v courses.VideoCore) Video {
	return Video{
		ID:       v.ID,
		Title:    v.Title,
		VideoID: v.VideoID,
		Duration: v.Duration,
	}
}
//


func (a *Video) toCore() courses.VideoCore {
	return courses.VideoCore{
		ID:          a.ID,
		Title:       a.Title,
		VideoID:     a.VideoID,
		Duration:    a.Duration,
	}
}
//func fromSliceVideoCore(v []courses.VideoCore) []Video {
//	//videos := make([]courses.VideoCore, len(v))
//	//
//	//for i, v := range v {
//	//	videos[i] = fromVideoCore(&v)
//	//}
//	//
//	//return videos
//	a := []courses.VideoCore{}
//	for key := range v {
//		a = append(a, v[key].toCore())
//	}
//	return a
//}
