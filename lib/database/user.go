package database

import (
	"github.com/NaufalHSyahputra/alterra-agmc/config"
	"github.com/NaufalHSyahputra/alterra-agmc/middlewares"
	"github.com/NaufalHSyahputra/alterra-agmc/models"
)

func GetUsers() (models.Users, error) {
	var users models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUser(id int) (models.User, error) {
	var user models.User

	if e := config.DB.First(&user, id).Error; e != nil {
		return user, e
	}
	return user, nil
}

func CreateUser(name, email, password string) (int, error) {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	if e := config.DB.Create(&user).Error; e != nil {
		return -1, e
	}
	return int(user.ID), nil
}

func UpdateUser(id int, name, email, password string) (int, error) {
	var user models.User
	if e := config.DB.First(&user, id).Error; e != nil {
		return -1, e
	}
	user.Name = name
	user.Email = email
	user.Password = password
	if e := config.DB.Save(&user).Error; e != nil {
		return -1, e
	}
	return int(user.ID), nil
}

func DeleteUser(id int) error {
	if e := config.DB.Delete(&models.User{}, id).Error; e != nil {
		return e
	}
	return nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.Users

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {

	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
