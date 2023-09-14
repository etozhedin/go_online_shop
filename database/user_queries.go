package database

import "online_shop/models"

func GetUsersList() ([]models.User, error) {
	var users []models.User

	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
