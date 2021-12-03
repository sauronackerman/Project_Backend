package routes

import (
 "RestfulAPIElearningVideo/factory"
 "github.com/labstack/echo/v4"
 "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo  {
 presenter := factory.New()

 e := echo.New()
 e.Pre(middleware.RemoveTrailingSlash())
 e.POST("/courses", presenter.CoursePresentation.CreateCourse)
 e.POST("/courses/:playlistId", presenter.CoursePresentation.AddVideoToCourse)
 e.POST("/:userId/uservideo/:videoId", presenter.UserPresentation.UserStartCourse)
 e.DELETE("/courses/:playlistId", presenter.CoursePresentation.DeleteCourse)
 e.DELETE("/tasks/:videoId", presenter.TaskPresentation.DeleteTask)
 e.POST("/:username/courses/:courseId", presenter.UserPresentation.PostUserCourse)
 e.POST("/login", presenter.UserPresentation.LoginUser)
 e.POST("/tasks", presenter.TaskPresentation.CreateTask)
 e.GET("/task/:videoId", presenter.TaskPresentation.GetTaskByVideoId)
 return e
}
