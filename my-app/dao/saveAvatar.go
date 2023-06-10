package dao

import (
	"my-app/model"
)

func SaveAvatar(username string, url string, hash string) error {
	var user model.User
	DB.Where("username = ?", username).First(&user)

	user.AvatarImage = url
	user.ImageHash = hash

	result := DB.Save(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetImageHash(username string) (string, error) {
	var user model.User
	result := DB.Where("username", username).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.ImageHash, nil

}
