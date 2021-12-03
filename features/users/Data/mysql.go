package Data

import (

	"RestfulAPIElearningVideo/features/users"
	"context"
	"errors"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gorm.io/gorm"
	"log"
)

type userData struct {
	Conn *gorm.DB
}

func NewUserData(conn *gorm.DB) users.Data {
	return &userData{
		conn,
	}
}



func (ud *userData) CheckUser(data users.UserCore) (users.UserCore, error) {
	var userData User

	err := ud.Conn.Where("username = ? and password = ?", data.Username, data.Password).First(&userData).Error

	if userData.Name == "" && userData.ID == 0 {
		return users.UserCore{}, errors.New("no existing user")
	}
	if err != nil {
		return users.UserCore{}, err
	}

	return toUserCore(userData), nil
}

func (ud *userData) SelectUserByUsername(username string) (users.UserCore, error) {
	user := User{}
	err := ud.Conn.Where("username = ?", username).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return users.UserCore{}, err
	}
	return toUserCore(user), nil
}
func (ud *userData) SelectUserById(userId uint) (users.UserCore, error) {
	user := User{}
	err := ud.Conn.Where("id = ?", userId).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return users.UserCore{}, err
	}
	return toUserCore(user), nil
}
func (ud *userData) SelectCourseByPlaylistId(playlistId string) (string, error) {
	course := users.UserCourse{}
	err := ud.Conn.Where("course_id = ?", course.CourseID).First(&course).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return course.CourseID, err
	}
	return playlistId, nil
}
//func (ud *userData) SelectTaskByVideoId(videoId string) (string, error) {
//	task := users.UserTaskCore{}
//	err := ud.Conn.Where("video_id = ?", task.TaskID).First(&task).Error
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		return task.UserTaskURL, err
//	}
//	return videoId, nil
//}


func (ud *userData) InsertUserCourse(courses users.UserCourse) error {
	newCourse := UserCourse{
		UserID:     courses.UserID,
		CourseID: courses.CourseID,
	}
	return ud.Conn.Create(&newCourse).Error
}

func (ud *userData)  GetPlaylistIdforVideo(ctx context.Context, playlistId string, userId uint) ([]users.UserCourseVideo, error)  {

	//API_KEY := os.Getenv("YT")
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey("AIzaSyDNbJBf7nypZKyj5SQFi_haZ66-SsNWIiM"))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	insert2 := youtubeService.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(playlistId).MaxResults(50)
	response, err := insert2.Do()
	if err != nil {
		panic(err)
	}
	var videos []UserCourseVideo
	var videoParse []string
	var video UserCourseVideo
	var videoparse string
	for _,item := range response.Items {
		//video.CourseID = playlistId
		//video.VideoID = item.ContentDetails.VideoId
		videoparse = item.ContentDetails.VideoId

		videoParse = append(videoParse, videoparse)
	}
	for i := 0; i < len(videoParse); i++ {
		insert3 := youtube.NewVideosService(youtubeService).List([]string{"snippet","contentDetails"}).Id(videoParse[i])
		resp, err := insert3.Do()
		if err != nil {
			panic(err)
		}
		for _,item := range resp.Items {
			video.UserID = userId
			video.CourseID = playlistId
			video.VideoID = videoParse[i]
			video.Title = item.Snippet.Title
			video.Duration = item.ContentDetails.Duration
			//video.Duration = (item.ContentDetails.Duration)
			videos = append(videos, video)
		}
	}

	err = ud.Conn.Create(videos).Error
	return toSliceUserVideoCore(videos), err
}

func (ud *userData) UpdateUserNoteData(id uint, data users.UserCourseVideo) (err error) {

	record := UserCourseVideo{
		Note:  data.Note,
	}
	if err = ud.Conn.Model(&record).Where("user_id = ?", id).Updates(UserCourseVideo{Note: record.Note}).Error; err != nil {
		return err
	}
	return err
}