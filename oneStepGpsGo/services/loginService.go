package services

import (
	"errors"
	"oneStepGps/entities"
	"oneStepGps/initializers"
	"oneStepGps/models"
)

func Login(apiKey string) (int, error) {
	var user entities.User
	initializers.DB.Where(&entities.User{ApiKey: apiKey}).First(&user)
	if user.ApiKey != "" {
		return user.Id, nil
	}

	if len(pingServer(apiKey)) == 0 {
		return -1, errors.New("key failed verification")
	}

	user = entities.User{ApiKey: apiKey}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return -1, errors.New("Can`t create object\n [ERROR] - " + result.Error.Error())
	}

	return user.Id, nil
}

func pingServer(apiKey string) []models.DeviceModel {
	pingResult := GetDevices(apiKey)
	return pingResult
}
