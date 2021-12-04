package mocks

import (
	"RestfulAPIElearningVideo/features/courses"
	"context"
	"github.com/stretchr/testify/mock"
)

type Data struct {
	mock.Mock
}

func (_m *Data) InsertCourse(course courses.CourseCore) (courses.CourseCore, error)  {
	ret := _m.Called(course)

	var r0 courses.CourseCore
	if rf, ok := ret.Get(0).(func(courses.CourseCore) courses.CourseCore); ok {
		r0 = rf(course)
	} else {
		r0 = ret.Get(0).(courses.CourseCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core courses.CourseCore) error); ok {
		r1 = rf(course)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Data) GetPlaylistIdforVideo(ctx context.Context, playlistId string) ([]courses.VideoCore, error) {
	ret := _m.Called(ctx, playlistId)

	var r0 []courses.VideoCore
	if rf, ok := ret.Get(0).(func(context.Context, string) []courses.VideoCore); ok {
		r0 = rf(ctx, playlistId)
	} else {
		r0 = ret.Get(0).([]courses.VideoCore)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx,playlistId)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *Data) SelectCourseById(id uint)  error {
	ret := _m.Called(id)
	var r1 error
	if rf, ok := ret.Get(1).(func( uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}
	return r1
}

func (_m *Data) SelectVideoByVideoId(videoId string)  error {
	ret := _m.Called(videoId)
	var r1 error
	if rf, ok := ret.Get(1).(func( string) error); ok {
		r1 = rf(videoId)
	} else {
		r1 = ret.Error(1)
	}
	return r1
}

func (_m *Data) DeleteCourseDataById(id string) (courses.CourseCore, error) {
	ret := _m.Called(id)

	var r0 courses.CourseCore

	if rf, ok := ret.Get(0).(func(string) courses.CourseCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(courses.CourseCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}