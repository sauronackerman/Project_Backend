package Presentation

import (
	"RestfulAPIElearningVideo/features/users"
	"RestfulAPIElearningVideo/features/users/Presentation/request"
	"RestfulAPIElearningVideo/features/users/Presentation/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UsersPresentation struct {
	usersBusiness users.Business
}
type json map[string]interface{}

func NewUserPresentation(up users.Business) *UsersPresentation  {
	return &UsersPresentation{usersBusiness: up}
}

func (up *UsersPresentation) LoginUser(c echo.Context) error  {
	user := request.UserAuth{}
	c.Bind(&user)
	data, err := up.usersBusiness.LoginUser(user.ToUserCore())
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserLoginResponse(data),
	})
}

func (up *UsersPresentation) PostUserCourse(c echo.Context) error {
	var username string
	var courseId string

	echo.PathParamsBinder(c).String("username", &username)
	echo.PathParamsBinder(c).String("courseId", &courseId)

	//issuer := c.Get("user").(jwt.MapClaims)

	//if issuer["username"] != username && issuer["role"] != "admin" {
	//	return c.JSON(http.StatusForbidden, json{
	//		"message": "Unauthorized user!",
	//	})
	//}

	err, status := up.usersBusiness.UserChooseCourse(courseId, username)
	if err != nil {
		return c.JSON(status, json{
			"message": "Failed get course",
			"error":   err.Error(),
		})
	}
	return c.JSON(status, json{
		"message": "Success get course",
	})
}

func (up *UsersPresentation)  UserStartCourse(c echo.Context) error {
	var userId string
	var videoId string

	echo.PathParamsBinder(c).String("userId", &userId)
	echo.PathParamsBinder(c).String("videoId", &videoId)
	var newVideo request.CreateVideo
	ctx := c.Request().Context()
	//req := c.Param("playlistId")
	userIdd,_ := strconv.Atoi(userId)
	c.Bind(&newVideo)
	video, err, status := up.usersBusiness.UserStartCourse(ctx, videoId, uint(userIdd))
	if err != nil {
		return c.JSON(status, json{
			"message": " ",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, json{"video": video})

}

func (up *UsersPresentation)  UpdateUserNoteData(c echo.Context) error {
	user := request.UserCourseVideo{}
	var userid string
	echo.PathParamsBinder(c).String("id", &userid)
	c.Bind(&user)
	err := up.usersBusiness.UpdateUserNote(user.UserID, request.ToNoteCoreVideo(user))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Note changed",
	})
}
