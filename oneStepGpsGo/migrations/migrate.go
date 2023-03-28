package main

import (
	"oneStepGps/entities"
	"oneStepGps/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(
		&entities.User{},
	)
}
