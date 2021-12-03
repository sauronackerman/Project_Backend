package Presentation

import (

	"RestfulAPIElearningVideo/features/courses"
	"RestfulAPIElearningVideo/features/courses/Presentation/request"
	"RestfulAPIElearningVideo/features/courses/Presentation/response"
	"github.com/labstack/echo/v4"
	"net/http"

)
type BaseReponses struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"data"`
}
//func NewErrorResponse(c echo.Context, error int, err error) error {
//	response := BaseReponses{}
//	response.Code = error
//	response.Message = "Failed"
//	response.Errors = []string{err.Error()}
//
//	return c.JSON(http.StatusBadRequest, response)
//}
//func NewSuccessResponse(c echo.Context, data interface{}) error {
//	response := BaseReponses{}
//	response.Code = http.StatusOK
//	response.Message = "Success"
//	response.Data = data
//	//log.Print(data)
//	return c.JSON(http.StatusOK, response)
//}
type CoursePresentation struct {
	courseBusiness courses.Business
}
type json map[string]interface{}
func NewPresentation(courseBusiness courses.Business) *CoursePresentation {
	return &CoursePresentation{courseBusiness}
}

func(cp *CoursePresentation) CreateCourse(c echo.Context) error {
	var newCourse request.CreateCourse
	c.Bind(&newCourse)
	course, err, status := cp.courseBusiness.CreateCourse(request.ToCore(newCourse))
	if err != nil {
		return c.JSON(status, json{
			"message": " ",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, json{"courses": course})
}



func(cp *CoursePresentation) AddVideoToCourse(c echo.Context) error {
	var newVideo request.CreateVideo
	ctx := c.Request().Context()
		req := c.Param("playlistId")
		c.Bind(&newVideo)
		video, err, status := cp.courseBusiness.AddVideoToCourse(ctx, req)
	if err != nil {
		return c.JSON(status, json{
			"message": " ",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, json{"video": video})

	}

func (cp *CoursePresentation) DeleteCourse(c echo.Context) error {
	playlistId := c.Param("playlistId")

	data, err := cp.courseBusiness.DeleteCourseById(playlistId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data deleted",
		"data":    response.ToCourseResponse(data)})
}

//	var newVideo request.CreateVideo
//	c.Bind(&newVideo)
//	course, err, status := cp.courseBusiness.AddVideoToCourse(request.ToCore(newCourse))
//	//ctx := c.Request().Context()
//	//requestParams := c.Param("playlistId")
//	//req := request.DetailCreateRequest{}
//	//err := c.Bind(&req)
//	//if err != nil {
//	//	return NewErrorResponse(c,http.StatusBadRequest,errors.New("something wrong with your request"))
//	//}
//	////serv domain
//	//log.Println("REQUEST PARAMS", requestParams)
//	//
//	//
//	//if err != nil {
//	//	return NewErrorResponse(c, http.StatusBadRequest,errors.New("error with your request params"))
//	//}
//	//addReq := request.CreateVideo{
//	//	ID:        0,
//	//	CreatedAt: time.Time{},
//	//	UpdatedAt: time.Time{},
//	//	Title:     "",
//	//	VideoID:   req.PlaylistID,
//	//	Duration:  0,
//	//}
//	//
//	//result, err := cp.courseBusiness.AddVideoToCourse(ctx,addReq)
//	//
//	//response := response.DetailCreateResponse{
//	//	PlaylistId: result.PlaylistId,
//	//	YoutubeId: result.YoutubeDataId,
//	//}
//	//
//	//return NewSuccessResponse(c, response)
//}

