package mocks

import (
	"RestfulAPIElearningVideo/features/courses"
	"context"
	"github.com/stretchr/testify/mock"
)

type Business struct {
	 mock.Mock
}

func (_m *Business) CreateCourse(course courses.CourseCore) (courses.CourseCore, error, int){
	ret := _m.Called(course)

	var m1 courses.CourseCore
	if rf, ok := ret.Get(0).(func(courses.CourseCore) courses.CourseCore); ok {
		m1 = rf(course)
	} else {
		m1 = ret.Get(0).(courses.CourseCore)
	}

	var m2 error
	if rf, ok := ret.Get(1).(func(courses.CourseCore) error); ok {
		m2 = rf(course)
	} else {
		m2 = ret.Error(1)
	}

	var m3 int
	if rf, ok := ret.Get(2).(func(courses.CourseCore) int); ok {
		m3 = rf(course)
	} else {
		m3 = ret.Get(2).(int)
	}

	return m1, m2, m3
}

func (_m *Business) AddVideoToCourse(ctx context.Context, playlistId string) ([]courses.VideoCore, error, int) {
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
	var r2 int
	if rf, ok := ret.Get(2).(func(context.Context, string) int); ok {
		r2 = rf(ctx, playlistId)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

func (_m *Business) DeleteCourseById(id string) (courses.CourseCore, error) {
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
	return r0,r1
}

func (_m *Business) FindVideoByVideoId(id string) (error, int) {
	ret := _m.Called(id)
	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}
	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(int)
	}
	return r0,r1
}

func (_m *Business) FindCourseById(id uint) (error, int) {
	ret := _m.Called(id)
	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}
	var r1 int
	if rf, ok := ret.Get(1).(func(uint) int); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(int)
	}
	return r0,r1
}