package Data

import (
	"RestfulAPIElearningVideo/features/users"
	//"RestfulAPIElearningVideo/migration"
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	Name string
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	UserCourses []UserCourse `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserCourseVideos []UserCourseVideo `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserTasks []UserTask `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	createdAt time.Time
	updatedAt time.Time
}

type UserCourse struct {
	ID       	uint `gorm:"primaryKey"`
	UserID uint
	CourseID string
	Videos []UserCourseVideo `gorm:"foreignKey:CourseID;type:longtext;reference:CourseID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	createdAt time.Time
	updatedAt time.Time
}

type UserCourseVideo struct {
	ID       	uint `gorm:"primaryKey"`
	UserID 		uint
	CourseID 	string
	Title 		string
	VideoID 	string `gorm:"not null"`
	Duration 	string
	//Task 		UserTask `gorm:"foreignKey:VideoID;type:longtext;reference:VideoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Note 		string `gorm:"default:'belum ada notes';constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UpdatedAt 	time.Time
	CreatedAt 	time.Time
}

type UserTask struct {
	ID       	uint `gorm:"primaryKey"`
	UserID uint
	VideoID string
	TaskURL string
	createdAt time.Time
	updatedAt time.Time
}



func 	toUserCore(u User) users.UserCore {
	return users.UserCore{
		ID:        u.ID,
		Name:      u.Name,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.createdAt,
		UpdatedAt: u.updatedAt,
	}
}
//func fromCore(c users.UserCore) User {
//
//	password, _ := migration.GenerateHashFromPass(c.Password)
//	return User{
//		Name:     c.Name,
//		Username:    c.Username,
//		Password: password,
//	}
//}
func toUserVideoCore(v *UserCourseVideo) users.UserCourseVideo {
	return users.UserCourseVideo{
		ID:       v.ID,
		Title:    v.Title,
		VideoID: v.VideoID,
		Duration: v.Duration,
	}
}
func toSliceUserVideoCore(v []UserCourseVideo) []users.UserCourseVideo {
	videos := make([]users.UserCourseVideo, len(v))

	for i, v := range v {
		videos[i] = toUserVideoCore(&v)
	}

	return videos
}

