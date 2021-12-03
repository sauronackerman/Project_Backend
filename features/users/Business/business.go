package Business

import (
	"RestfulAPIElearningVideo/features/users"
	"RestfulAPIElearningVideo/middleware"
	"context"
	"errors"
	"net/http"

	//"RestfulAPIElearningVideo/migration"
)

type usersBusiness struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business  {
	return &usersBusiness{usrData}
}

func (ub *usersBusiness) LoginUser(data users.UserCore) (users.UserCore, error) {
	userData,err:=ub.userData.CheckUser(data)
	if err != nil {
		return users.UserCore{}, err
	}

	userData.Token, err = middleware.CreateTokens(userData.ID, userData.Name)
	if err != nil {
		return userData, err
	}
	return userData, err
}

func (ub *usersBusiness) UserChooseCourse(username string, playlistId string) (error, int) {
	user, err := ub.userData.SelectUserByUsername(username)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	if user.IsNotFound() {
		return errors.New("User not found"), http.StatusNotFound
	}

	playlistId, err = ub.userData.SelectCourseByPlaylistId(playlistId)
	if err != nil {
		return err, http.StatusInternalServerError
	}


	courses := users.UserCourse{
		UserID:     user.ID,
		CourseID: playlistId,
	}
	err = ub.userData.InsertUserCourse(courses)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusCreated
}


func (ub *usersBusiness) UserStartCourse(ctx context.Context, playlistId string, userId uint) ([]users.UserCourseVideo, error, int) {
	createdVideo, err := ub.userData.GetPlaylistIdforVideo(ctx, playlistId, userId)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	user, err := ub.userData.SelectUserById(userId)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	if user.IsNotFound() {
		return nil, errors.New("User not found"), http.StatusNotFound
	}
	//err = ub.userData.InsertUserVideo(createdVideo)
	//if err != nil {
	//	return nil,err, http.StatusInternalServerError
	//}
	//return users.UserCourseVideo{},nil, http.StatusCreated
	return createdVideo, nil, http.StatusOK
}

func (ub *usersBusiness) UpdateUserNote(id uint, data users.UserCourseVideo) error {
	err := ub.UpdateUserNote(id, data)
	if err != nil {
		return err
	}
	return nil
}

//func (ub *usersBusiness) UserInsertTask(username string, videoId string) (error, int) {
//	user, err := ub.userData.SelectUserByUsername(username)
//	if err != nil {
//		return err, http.StatusInternalServerError
//	}
//	if user.IsNotFound() {
//		return errors.New("User not found"), http.StatusNotFound
//	}
//
//	videoId, err = ub.userData.SelectTaskByVideoId(videoId)
//	if err != nil {
//		return err, http.StatusInternalServerError
//	}
//
//
//	task := users.UserTaskCore{
//		UserID:     user.ID,
//		TaskID: videoId,
//	}
//	err = ub.userData.InsertUserTask(task)
//	if err != nil {
//		return err, http.StatusInternalServerError
//	}
//	return nil, http.StatusCreated
//}
//func (cb *courseBusiness) AddVideoToCourse(ctx context.Context, playlistId string) ([]courses.VideoCore, error, int) {
//	createdVideo, err := cb.courseData.GetPlaylistIdforVideo(ctx, playlistId)
//	if err != nil {
//		return nil, err, http.StatusInternalServerError
//	}
//	return createdVideo, nil, http.StatusOK
//}

//
//func (cb *courseBusiness) FindCourseById(id uint) (error, int) {
//	err := cb.courseData.SelectCourseById(id)
//	if err != nil {
//		return err, http.StatusInternalServerError
//	}
//
//	//userData, err, _ := cb.userBusiness.FindUserById(articleData.AuthorID)
//	//if err != nil {
//	//	return articles.ArticleCore{}, err, http.StatusInternalServerError
//	//}
//	//
//	//articleData.Author.Username = userData.Username
//	//articleData.Author.Email = userData.Email
//	//articleData.Author.Name = userData.Name
//
//	return nil, http.StatusOK
//}
//
//func (cb *courseBusiness) DeleteCourseById(id uint) (error, int) {
//	err := cb.courseData.DeleteCourseDataById(id)
//	if err != nil {
//		return err,	 http.StatusInternalServerError
//	}
//
//	return nil, http.StatusAccepted
//}
//
