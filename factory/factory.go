package factory

import (
	"PROJECT_BACKEND/config"
	courseBusiness "PROJECT_BACKEND/features/courses/Business"
	courseData "PROJECT_BACKEND/features/courses/Data"
	coursePresentation "PROJECT_BACKEND/features/courses/Presentation"
	taskBusiness "PROJECT_BACKEND/features/tasks/Business"
	taskData "PROJECT_BACKEND/features/tasks/Data"
	taskPresentation "PROJECT_BACKEND/features/tasks/Presentation"
	userBusiness "PROJECT_BACKEND/features/users/Business"
	userData "PROJECT_BACKEND/features/users/Data"
	userPresentation "PROJECT_BACKEND/features/users/Presentation"
)

type Presenter struct {
	CoursePresentation *coursePresentation.CoursePresentation
	UserPresentation   *userPresentation.UsersPresentation
	TaskPresentation   *taskPresentation.TaskPresentation
}

func New() *Presenter {
	//courses
	courseData := courseData.NewCourseData(config.DB)
	courseBusiness := courseBusiness.NewCourseBusiness(courseData)
	coursePresentation := coursePresentation.NewPresentation(courseBusiness)

	//users
	userData := userData.NewUserData(config.DB)
	userBusiness := userBusiness.NewUserBusiness(userData)
	userPresentation := userPresentation.NewUserPresentation(userBusiness)

	//task
	taskData := taskData.NewTaskData(config.DB)
	taskBusiness := taskBusiness.NewTaskBusiness(taskData, courseBusiness)
	taskPresentation := taskPresentation.NewTaskPresentation(taskBusiness)

	return &Presenter{
		coursePresentation,
		userPresentation,
		taskPresentation,
	}
}
