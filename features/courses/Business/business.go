package Business

import (
	"RestfulAPIElearningVideo/features/courses"
	"context"
	"net/http"
)

type courseBusiness struct {
	courseData courses.Data
}

func NewCourseBusiness(crsData courses.Data) courses.Business  {
	return &courseBusiness{crsData}
}

func (cb *courseBusiness) CreateCourse(course courses.CourseCore) (courses.CourseCore, error, int) {
	createdCourse, err := cb.courseData.InsertCourse(course)
	if err != nil {
		return course, err, http.StatusInternalServerError
	}
	return createdCourse, nil, http.StatusOK
}

func (cb *courseBusiness) AddVideoToCourse(ctx context.Context, playlistId string) ([]courses.VideoCore, error, int) {
	createdVideo, err := cb.courseData.GetPlaylistIdforVideo(ctx, playlistId)
	if err != nil {
		return []courses.VideoCore{}, err, http.StatusInternalServerError
	}
	return createdVideo, nil, http.StatusOK
}


func (cb *courseBusiness) FindCourseById(id uint) (error, int) {
	err := cb.courseData.SelectCourseById(id)
		if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
func (cb *courseBusiness) FindVideoByVideoId(videoId string) (error, int) {
	err := cb.courseData.SelectVideoByVideoId(videoId)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (cb *courseBusiness) DeleteCourseById(id string) (courses.CourseCore, error) {
	data, err := cb.courseData.DeleteCourseDataById(id)
	if err != nil {
		return courses.CourseCore{}, err
	}

	return data, nil
}
