package dao

import (
	"log"
	"my-app/logs"
	"my-app/model"
)

func IsExist(username string) bool {
	var user model.User
	result := DB.First(&user, "username = ?", username)

	if result.Error != nil {
		logs.LogError(result.Error)
		return false
	}

	return true
}

func AddUser(user model.User) error {

	result := DB.Create(&user)

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	log.Println("add user successfully")
	return nil

}

func QueryPassword(username string) (string, error) {
	var user model.User

	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		logs.LogError(result.Error)
		return "", result.Error
	}

	return user.Password, nil

}

func GetUserDetails(username string) (model.User, error) {
	var user model.User
	result := DB.Where("username = ?", username).Find(&user)

	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func SaveUserDetails(username string, telephone string, address string) error {
	var user model.User
	DB.Where("username = ?", username).Find(&user)
	user.Telephone = telephone
	user.Address = address
	result := DB.Where("username = ?", username).Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
