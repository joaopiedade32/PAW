package service

import (
	"bookapi/entity"
	"errors"
)

var Users []entity.User

func InitializeUsers() {
	Users = []entity.User{}
}

func GetAllUsers() []entity.User {
	return Users
}

func GetUserProfile(userId uint64) (entity.User, error) {
	for i := 0; i < len(Users); i++ {
		if Users[i].ID == userId {
			return Users[i], nil
		}
	}
	return entity.User{}, errors.New("user dont exist")
}

func RegisterUser(user entity.User) entity.User {
	user.ID = uint64(len(Users) + 1)
	Users = append(Users, user)
	return user
}

func UpdateUserProfile(user entity.User) (entity.User, error) {
	for i := 0; i < len(Users); i++ {
		if Users[i].ID == user.ID {
			Users[i] = user
			return user, nil
		}
	}
	return user, errors.New("user dont exist")
}

func DeleteUser(userId uint64) error {
	for i := 0; i < len(Users); i++ {
		if Users[i].ID == userId {
			Users = append(Users[:i], Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user dont exist")
}
