package dao

import (
	"log"
	"my-app/model"
)

func IsExist(username string) bool {
	var user model.User

	//sqlStr := "select username,password from users where username = ?"

	result := DB.First(&user, "username = ?", username)

	if result.Error != nil {
		log.Println(result.Error)
		return false
	}

	return true
}

func AddUser(user model.User) error {

	//sqlStr := "insert into users (username, password) values (?,?)"

	result := DB.Create(&user)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	log.Println("add user successfully")
	return nil

}

func QueryPassword(username string) (string, error) {

	//var password string
	//sqlStr := "select password from users where username = ?"
	var user model.User

	result := DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		log.Println(result.Error)
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
