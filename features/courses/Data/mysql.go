package Data

import (
	"RestfulAPIElearningVideo/features/courses"
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gorm.io/gorm"
	"log"
)

type courseData struct {
	Conn *gorm.DB
}

func NewCourseData(conn *gorm.DB) *courseData  {
	return &courseData{conn}
}

func (cd *courseData) InsertCourse(course courses.CourseCore)  (resp courses.CourseCore, err error) {
	//record := toCourseCore(course)
	videos := make([]Video, len(course.Videos))
	for i, v := range course.Videos {
		videos[i] = Video{VideoID: v.VideoID}
	}
	newCourse := Course{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		PlaylistID:  course.PlaylistID,
		Videos:      videos,
	}
	//if err := cd.Conn.Create(&record).Error; err != nil {
	//	return courses.CourseCore{}, err
	//}
	//
	//return course, nil

	err = cd.Conn.Create(&newCourse).Error
	return toCourseCore(&newCourse), err
}
func (cd *courseData) GetPlaylistIdforVideo(ctx context.Context, playlistId string) ([]courses.VideoCore, error)  {

	//API_KEY := os.Getenv("YT")
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey("AIzaSyDNbJBf7nypZKyj5SQFi_haZ66-SsNWIiM"))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
	//insert := youtubeService.Search.List([]string{"playlistId","contentDetails"}).Q(playlistId)
	insert2 := youtubeService.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(playlistId).MaxResults(50)
	response, err := insert2.Do()
	if err != nil {
		panic(err)
	}
	var videos []Video
	var videoParse []string
	var video Video
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
			video.CourseID = playlistId
			video.VideoID = videoParse[i]
			video.Title = item.Snippet.Title
			video.Duration = item.ContentDetails.Duration
			//video.Duration = (item.ContentDetails.Duration)
			videos = append(videos, video)
		}
	}

	err = cd.Conn.Create(videos).Error
	return toSliceVideoCore(videos), err
}
func (cd *courseData) SelectCourseById(id uint) error {
	course := Course{}
	err := cd.Conn.First(&course, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return  err
	}
	return nil
}
func (cd *courseData) SelectVideoByVideoId(videoId string) error {
	video := Video{}
	err := cd.Conn.First(&video, videoId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return  err
	}
	return nil
}


func (cd *courseData) DeleteCourseDataById(p string) (courses.CourseCore, error) {
	var course Course
	err := cd.Conn.Where("playlistId = ?", p).Delete(&course).Error
	if err != nil {
		return courses.CourseCore{}, err
	}
	return toCore2(course), nil
}
