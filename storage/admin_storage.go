package storage

import (
	"errors"
	"postmanapi/models"

	"github.com/jinzhu/gorm"
)

var users []models.User

func GetAllUsers() []models.User {
	return users
}

func GetUserByID(id int) *models.User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func AddUser(user models.User) {
	users = append(users, user)
}

func UpdateUser(id int, updatedUser models.User) {
	for index, user := range users {
		if user.ID == id {
			users[index] = updatedUser
			break
		}
	}
}

func DeleteUser(id int) {
	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
}

func GetUserByUsernameAndPassword(username, password string) (*models.Admin, error) {
	var user models.Admin
	err := DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
