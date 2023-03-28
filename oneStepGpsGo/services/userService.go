package services

import (
	"errors"
	"oneStepGps/entities"
	"oneStepGps/initializers"
)

func GetProfileData(Id int) entities.User {
	var user entities.User
	initializers.DB.First(&user, Id)
	return user
}

func UpdateUser(user entities.User) error {
	var u entities.User
	initializers.DB.First(&u, user.Id)
	if u.Id == 0 {
		return errors.New("User not found")
	}

	u = user
	result := initializers.DB.Save(&u)
	return result.Error
}
