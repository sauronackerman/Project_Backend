package main

import (
	"RestfulAPIElearningVideo/database"
	"RestfulAPIElearningVideo/middleware"
	"RestfulAPIElearningVideo/migration"
	"RestfulAPIElearningVideo/routes"
)

func main() {
	database.InitDB()
	migration.AutoMigrate()
	e := routes.New()

	e.Start(":1234")
	middleware.LogMiddleware(e)
}


