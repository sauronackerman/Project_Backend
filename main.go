package main

import (
	"RestfulAPIElearningVideo/config"
	"RestfulAPIElearningVideo/middleware"
	"RestfulAPIElearningVideo/migration"
	"RestfulAPIElearningVideo/routes"
)

func main() {
	config.InitDB()
	migration.AutoMigrate()
	e := routes.New()

	e.Start(":1234")
	middleware.LogMiddleware(e)
}

