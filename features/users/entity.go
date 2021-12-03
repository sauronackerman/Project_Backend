package users

import (
	"context"
	"time"
)

type UserCore struct {
	ID uint
	Name     	string
	Username   	string
	Password 	string
	UserCourses []UserCourse
	UserTasks []UserTaskCore
	Token       string
	CreatedAt time.Time
	UpdatedAt time.Time

}

type UserLog struct {
	ID int
	Name string
	Username string
	Token string
}

type UserTaskCore struct {
	ID uint
	UserID uint
	TaskID string
	UserTaskURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}




type UserCourse struct {
	ID uint
	CourseID  string
	UserID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCourseVideo struct {
	ID       	uint `gorm:"primaryKey"`
	UserID 		uint
	CourseID 	string
	Title 		string
	VideoID 	string `gorm:"not null"`
	Note 		string
	Duration 	string
	UpdatedAt 	time.Time
	CreatedAt 	time.Time
}

type Business interface {
	LoginUser(data UserCore) (UserCore, error)
	UserChooseCourse(username string, playlistId string) (error, int)
	UserStartCourse(ctx context.Context, playlistId string, userId uint) ([]UserCourseVideo, error, int)
	UpdateUserNote(id uint, data UserCourseVideo) error
	//UserInsertTask(username string, videoId string) (error, int)
	//FindUserByUsername(username string) (UserCore, error, int)
	////FindUsers() ([]UserCore, error, int)
	//FindUserCourses(username string) (UserCore, error, int)
	//FindUserTask(username string) (UserCore, error, int)
	//FindUserNote(username string) (UserCore, error, int)
	////FindUsersByIds(ids []uint) ([]UserCore, error, int)
	////FindUserById(id uint) (UserCore, error, int)
	//
	////FindUserByEmail(email string) (UserCore, error, int)
	////CreateUser(user UserCore) (UserCore, error, int)
	//CreateUsers(username string, followerUsername string) (error, int)
	////EditUser(user UserCore) (UserCore, error, int)
	////RemoveUser(username string) (error, int)
	//RemoveCourse(username string, followingUsername string) (error, int)
}

type Data interface {
	//SelectUserByUsername(username string) (user UserCore,err error)
	CheckUser(data UserCore) (UserCore, error)
	SelectUserByUsername(username string) (UserCore, error)
	SelectUserById(userId uint) (UserCore, error)
	SelectCourseByPlaylistId(playlistId string) (string, error)
	GetPlaylistIdforVideo(ctx context.Context, playlistId string, userId uint) ([]UserCourseVideo, error)
	//SelectTaskByVideoId(videoId string) (string, error)
	InsertUserCourse(courses UserCourse) error
	UpdateUserNoteData(id uint, data UserCourseVideo) (err error)
	//InsertUserVideo(videos []UserCourseVideo) error
	//InsertUserTask(tasks UserTaskCore) error
	//SelectUserFollowers(userID uint) ([]FollowerCore, error)
	//SelectUserFollowings(userID uint) ([]FollowerCore, error)
	//SelectUserById(id uint) (UserCore, error)
	//SelectUsersByIds(ids []uint) ([]UserCore, error)
	//SelectUserByEmail(email string) (UserCore, error)
	//InsertUser(user UserCore) (UserCore, error)
	//InsertFollower(follower FollowerCore) error
	//UpdateUser(user UserCore) (UserCore, error)
	//DeleteUser(username string) error
	//DeleteFollowing(following FollowerCore) error
	//DeleteAllUserFollow(userID uint) error
}

func (u *UserCore) IsNotFound() bool {
	return u.ID == 0
}
//repo = business
