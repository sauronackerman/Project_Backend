package request

import (
	"RestfulAPIElearningVideo/features/users"
)

type UserCourseVideo struct {
	UserID 		uint
	//CourseID 	string
	//Title 		string
	//VideoID 	string `gorm:"not null"`
	//Duration 	string
	//Task 		UserTask `gorm:"foreignKey:VideoID;type:longtext;reference:VideoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Note 		string `gorm:"default:'belum ada notes';constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//UpdatedAt 	time.Time
	//CreatedAt 	time.Time
}

func ToNoteCoreVideo(req UserCourseVideo) users.UserCourseVideo {
	return users.UserCourseVideo{
		UserID:       req.UserID,
		Note: req.Note,
	}
}

type CreateVideo struct {
	//ID uint `json:"id"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
	//Title string `json:"title"`
	UserId uint `json:"user_id"`
	VideoID string `json:"video_id"`
	//Duration int `json:"duration"`
}


type UserAuth struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (data *UserAuth) ToUserCore() users.UserCore {
	return users.UserCore{
		Username:    data.Username,
		Password: data.Password,
	}
}
//func ToUserCore(req User) users.UserCore  {
//	return users.UserCore{
//		ID:          req.ID,
//		Name:        req.Name,
//		Username:    req.Username,
//		Password:    req.Password,
//	}
//}