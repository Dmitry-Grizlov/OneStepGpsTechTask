package main

import (
	"log"
	"oneStepGps/entities"
	"oneStepGps/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
}

func main() {
	db, err := initializers.CreateDbConnection()
	if err != nil {
		log.Fatal("Unexpected error occured while creating database connection.\n[ERROR] -", err.Error())
	}

	db.AutoMigrate(&entities.User{})
}
