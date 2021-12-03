package Business

import (
	"RestfulAPIElearningVideo/features/tasks"
	"github.com/stretchr/testify/mock"
)

type Data struct {
	mock.Mock
}

func (_m *Data) InsertData(data tasks.TaskCore) (tasks.TaskCore, error) {
	ret := _m.Called(data)
	var r0 tasks.TaskCore
	if rf, ok := ret.Get(0).(func(tasks.TaskCore) tasks.TaskCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(tasks.TaskCore)
	}
	var m2 error
	if rf, ok := ret.Get(1).(func(tasks.TaskCore) error); ok {
		m2 = rf(data)
	} else {
		m2 = ret.Error(1)
	}

	return r0, m2
}


func (_m *Data) SelectTaskByVideoId(videoId string) (tasks.TaskCore, error){
	ret := _m.Called(videoId)
	var r0 tasks.TaskCore
	if rf, ok := ret.Get(0).(func(string) tasks.TaskCore); ok {
		r0 = rf(videoId)
	} else {
		r0 = ret.Get(0).(tasks.TaskCore)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(videoId)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}


func (_m *Data) DeleteTask(videoId string) (tasks.TaskCore, error) {
	ret := _m.Called(videoId)

	var r0 tasks.TaskCore
	if rf, ok := ret.Get(0).(func(string) tasks.TaskCore); ok {
		r0 = rf(videoId)
	} else {
		r0 = ret.Get(0).(tasks.TaskCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(videoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}




