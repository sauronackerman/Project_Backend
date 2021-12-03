package migration

import (
	"PROJECT_BACKEND/config"
	course "PROJECT_BACKEND/features/courses/Data"
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

func AutoMigrate() {
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
		// courses
		&course.Course{},
		&course.Video{},
	)
	if err != nil {
		panic(err)
	}

}
