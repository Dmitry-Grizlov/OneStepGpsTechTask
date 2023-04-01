package services

import (
	"errors"
	"log"
	"oneStepGps/entities"
	"oneStepGps/initializers"
	"oneStepGps/models"
)

func Login(apiKey string) (int, error) {
	var user entities.User
	db, err := initializers.CreateDbConnection()
	if err != nil {
		log.Println("Unexpected error occured while creating database connection.\n[ERROR] -", err.Error())
		return -1, err
	}

	db.Where(&entities.User{ApiKey: apiKey}).First(&user)

	if user.ApiKey != "" {
		return user.Id, nil
	}

	if len(pingServer(apiKey)) == 0 {
		return -1, errors.New("key verification failed")
	}

	user = entities.User{ApiKey: apiKey}
	result := db.Create(&user)
	if result.Error != nil {
		return -1, errors.New("Error creating object\n [ERROR] - " + result.Error.Error())
	}

	return user.Id, nil
}

func pingServer(apiKey string) []models.DeviceModel {
	pingResult, err := GetDevices(apiKey)
	if err != nil {
		return []models.DeviceModel{}
	}

	return pingResult
}
