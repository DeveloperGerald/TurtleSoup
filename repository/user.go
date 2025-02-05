package repository

import (
	"fmt"

	"github.com/DeveloperGerald/TurtleSoup/model"
	"github.com/google/uuid"
)

func GetUserByID(id string) (*model.User, error) {
	var users []model.User
	err := db.Where("id = ?", id).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("get user by id %s error: %v", id, err)
	}

	if len(users) > 1 {
		return nil, fmt.Errorf("get more than one user by id %s", id)
	}
	if len(users) < 1 {
		return nil, fmt.Errorf("find no user by id %s", id)
	}

	user := users[0]
	return &user, nil
}

func GetUserByName(name string) (*model.User, error) {
	var users []model.User
	err := db.Where("name = ?", name).Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("get user by name %s error: %v", name, err)
	}

	if len(users) > 1 {
		return nil, fmt.Errorf("get more than one user by name %s", name)
	}
	if len(users) < 1 {
		return nil, fmt.Errorf("find no user by name %s", name)
	}

	user := users[0]
	return &user, nil
}

func CreateUser(user model.User) (*model.User, error) {
	if user.ID == "" {
		newUUID := uuid.New()
		user.ID = newUUID.String()
	}
	err := db.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("create user error: %v", err)
	}

	return &user, nil
}

func DeleteUser(id string) error {
	var user model.User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return fmt.Errorf("delete user(%s) error: %v", id, err)
	}

	return nil
}
