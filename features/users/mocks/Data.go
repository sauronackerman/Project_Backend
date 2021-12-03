package mocks

import (
	"RestfulAPIElearningVideo/features/users"
	"context"
	"github.com/stretchr/testify/mock"
)

type Data struct {
	mock.Mock
}

func (_m *Data) CheckUser(data users.UserCore) (users.UserCore, error) {
	ret := _m.Called(data)

	var r0 users.UserCore
	if rf, ok := ret.Get(0).(func(users.UserCore) users.UserCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Data) SelectUserByUsername(username string) (users.UserCore, error) {
	ret := _m.Called(username)

	var r0 users.UserCore
	if rf, ok := ret.Get(0).(func(string) users.UserCore); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Data) SelectUserById(userId uint) (users.UserCore, error) {
	ret := _m.Called(userId)
	var r0 users.UserCore
	if rf, ok := ret.Get(0).(func(uint) users.UserCore); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(users.UserCore)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Data) SelectCourseByPlaylistId(playlistId string) (string, error) {
	ret := _m.Called(playlistId)
	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(playlistId)
	} else {
		r0 = ret.Get(0).(string)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(playlistId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Data) GetPlaylistIdforVideo(ctx context.Context, playlistId string, userId uint) ([]users.UserCourseVideo, error) {
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
	return r0, r1
}

func (_m *Data) InsertUserCourse(courses users.UserCourse) error {
	ret := _m.Called(courses)


	var m2 error
	if rf, ok := ret.Get(0).(func( course users.UserCourse) error); ok {
		m2 = rf(courses)
	} else {
		m2 = ret.Error(0)
	}

	return m2
}



func (_m *Data) UpdateUserNoteData(id uint, data users.UserCourseVideo) error {
	ret := _m.Called(id, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.UserCourseVideo) error); ok {
		r0 = rf(id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
