package dao

import "my-app/model"

func SaveAvatar(username string, url string) error {
	result := DB.Model(&model.User{}).Where("username = ?", username).Update("avatar_image", url)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
