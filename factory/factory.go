package factory

import (

	"RestfulAPIElearningVideo/config"
	courseBusiness "RestfulAPIElearningVideo/features/courses/Business"
	courseData "RestfulAPIElearningVideo/features/courses/Data"
	taskBusiness "RestfulAPIElearningVideo/features/tasks/Business"
	taskData "RestfulAPIElearningVideo/features/tasks/Data"
	taskPresentation "RestfulAPIElearningVideo/features/tasks/Presentation"
	userBusiness "RestfulAPIElearningVideo/features/users/Business"
	userData "RestfulAPIElearningVideo/features/users/Data"
	userPresentation "RestfulAPIElearningVideo/features/users/Presentation"

	coursePresentation "RestfulAPIElearningVideo/features/courses/Presentation"

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
