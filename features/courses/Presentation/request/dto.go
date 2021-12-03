package request

import (
	"RestfulAPIElearningVideo/features/courses"
)

type CreateCourse struct {
	//ID        uint `json:"id"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
	//DeletedAt sql.NullTime `json:"deleted_at"`
	Title string `json:"title"`
	Description string `json:"description"`
	PlaylistID string `json:"playlist_id"`
	Video []CreateVideo `json:"video"`
}

type CreateVideo struct {
	//ID uint `json:"id"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
	//Title string `json:"title"`
	VideoID string `json:"video_id"`
	//Duration int `json:"duration"`
}


//func (req *CreateesVideo) toVideoCore() courses.VideoCore  {
//	return courses.VideoCore{
//		ID:        req.ID,
//		Title:     req.Title,Video
//		VideoID:  req.VideoID,
//		Duration:  req.Duration,
//	}
//}
//



func ToCore(req CreateCourse) courses.CourseCore {
	return courses.CourseCore{
		//ID:          req.ID,
		//CreatedAt:   req.CreatedAt,
		//UpdatedAt:   req.UpdatedAt,
		Title:       req.Title,
		Description: req.Description,
		PlaylistID:   req.PlaylistID,
		//Videos: ToVideoCore(req.Video),
	}
}

//func ToCore2(req CreateCourse) courses.CourseCore {
//	return courses.CourseCore{
//		//ID:          req.ID,
//		//CreatedAt:   req.CreatedAt,
//		//UpdatedAt:   req.UpdatedAt,
//		//Title:       req.Title,
//		//Description: req.Description,
//		//PlaylistID:   req.PlaylistID,
//		Videos: req.Video,
//	}
//}





type GetCourseInput struct {
	ID int `uri:"id" binding:"required"`
}

type DetailCreateRequest struct {
	PlaylistID string `json:"playlist_id"`
}