package mocks

import (
	"RestfulAPIElearningVideo/features/users"
	"context"
	"github.com/stretchr/testify/mock"
)

type Business struct {
	mock.Mock
}

func (_m *Business) LoginUser(data users.UserCore) (users.UserCore, error)  {
	ret := _m.Called(data)

	var m1 users.UserCore
	if rf, ok := ret.Get(0).(func(users.UserCore) users.UserCore); ok {
		m1 = rf(data)
	} else {
		m1 = ret.Get(0).(users.UserCore)
	}

	var m2 error
	if rf, ok := ret.Get(1).(func( users.UserCore) error); ok {
		m2 = rf(data)
	} else {
		m2 = ret.Error(1)
	}

	return m1, m2
}

func (_m *Business) UserChooseCourse(username string, playlistId string) (error, int) {
	ret := _m.Called(username, playlistId)

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username,playlistId)
	} else {
		r1 = ret.Error(1)
	}
	var r2 int
	if rf, ok := ret.Get(2).(func(string, string) int); ok {
		r2 = rf(username, playlistId)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r1, r2
}

func (_m *Business) UserStartCourse(ctx context.Context, playlistId string, userId uint) ([]users.UserCourseVideo, error, int) {
	ret := _m.Called(ctx, playlistId, userId)

	var r0 []users.UserCourseVideo
	if rf, ok := ret.Get(0).(func(context.Context, string, uint) []users.UserCourseVideo); ok {
		r0 = rf(ctx, playlistId, userId)
	} else {
		r0 = ret.Get(0).([]users.UserCourseVideo)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint) error); ok {
		r1 = rf(ctx,playlistId, userId)
	} else {
		r1 = ret.Error(1)
	}
	var r2 int
	if rf, ok := ret.Get(2).(func(context.Context, string, uint) int); ok {
		r2 = rf(ctx, playlistId, userId)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

func (_m *Business) UpdateUserNoteData(id uint, data users.UserCourseVideo) error {
	ret := _m.Called(id, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.UserCourseVideo) error); ok {
		r0 = rf(id,data)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
