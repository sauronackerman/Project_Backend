package main

import (
	"RestfulAPIElearningVideo/config"
	"RestfulAPIElearningVideo/migration"
	"RestfulAPIElearningVideo/routes"
)

func main() {
	config.InitDB()
	migration.AutoMigrate()
	e := routes.New()
	e.Start(":1234")
}

