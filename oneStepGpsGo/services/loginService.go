package services

import (
	"errors"
	"oneStepGps/entities"
	"oneStepGps/initializers"
	"oneStepGps/models"
	"oneStepGps/utils"
)

func Login(apiKey string) (models.LoginModel, error) {
	var user entities.User

	db, err := initializers.CreateDbConnection()
	if err != nil {
		utils.ErrorCreateDbConnection(err)
		return models.LoginModel{Id: -1}, err
	}

	db.Where(&entities.User{ApiKey: apiKey}).First(&user)

	if user.ApiKey != "" {
		return models.LoginModel{Id: user.Id, ApiKey: apiKey}, nil
	}

	if len(pingServer(apiKey)) == 0 {
		return models.LoginModel{Id: -1}, errors.New("key verification failed")
	}

	user = entities.User{ApiKey: apiKey}
	result := db.Create(&user)
	if result.Error != nil {
		return models.LoginModel{Id: -1}, errors.New("Error creating object\n [ERROR] - " + result.Error.Error())
	}

	return models.LoginModel{Id: user.Id, ApiKey: apiKey}, nil
}
