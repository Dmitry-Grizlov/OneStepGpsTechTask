package services

import (
	"errors"
	"log"
	"oneStepGps/entities"
	"oneStepGps/initializers"
	"oneStepGps/utils"
)

func GetProfileData(Id int) (entities.User, error) {
	var user entities.User
	db, err := initializers.CreateDbConnection()
	if err != nil {
		log.Println("Unexpected error occured while creating database connection.\n[ERROR] -", err.Error())
		return user, err
	}

	db.First(&user, Id)
	if user.Id == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

func UpdateUser(user entities.User) error {
	var u entities.User
	db, err := initializers.CreateDbConnection()
	if err != nil {
		utils.ErrorCreateDbConnection(err)
		return err
	}

	db.First(&u, user.Id)
	if u.Id == 0 {
		return errors.New("user not found")
	}

	if u == user {
		return nil
	}

	if u.ApiKey != user.ApiKey && len(pingServer(user.ApiKey)) == 0 {
		return errors.New("new key is invalid")
	}

	u = user
	result := db.Save(&u)
	return result.Error
}
