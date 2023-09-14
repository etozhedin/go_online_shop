package database

import "online_shop/models"

func GetUsersNicknames() ([]string, error) {
	var nicknames []string

	result := DB.Model(&models.User{}).Select("username").Find(&nicknames)
	if result.Error != nil {
		return nil, result.Error
	}
	return nicknames, nil
}
