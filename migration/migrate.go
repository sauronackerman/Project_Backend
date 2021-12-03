package migration

import (
	"RestfulAPIElearningVideo/config"
	course "RestfulAPIElearningVideo/features/courses/Data"
	tasks "RestfulAPIElearningVideo/features/tasks/Data"
	user "RestfulAPIElearningVideo/features/users/Data"

	//"golang.org/x/crypto/bcrypt"
)

//func GenerateHashFromPass(password string) (string, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//	return string(bytes), err
//}
//
//func CompareHashAndPass(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}UpdatedAt 	time.Time
//	CreatedAt 	time.Time



func AutoMigrate()  {
	db := config.DB
	if err := db.Exec("DROP TABLE IF EXISTS user_course_videos").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS user_courses").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS user_notes").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS user_tasks").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS tasks").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS videos").Error; err != nil {
		panic(err)
	}
	if err := db.Exec("DROP TABLE IF EXISTS courses").Error; err != nil {
		panic(err)
	}

	err := db.AutoMigrate(
		&course.Course{},
		&course.Video{},
		&tasks.Task{},
		&user.User{},
		&user.UserTask{},
		&user.UserCourse{},
		//&user.UserNote{},
		&user.UserCourseVideo{},
	)
	if err != nil {
		panic(err)
	}
	//pass1, _ := GenerateHashFromPass("coba")
	user1 := user.User{
		Name:     "Mahard",
		Username:    "mahard",
		Password: "123",
	}

	//pass2, _ := GenerateHashFromPass("coba123")
	user2 := user.User{
		Name:     "Riza",
		Username:    "riza",
		Password: "coba133",
	}

	if err := db.Create(&user1).Error;
	err != nil {
		panic(err)
	}
	if err := db.Create(&user2).Error;
		err != nil {
		panic(err)
	}
}
