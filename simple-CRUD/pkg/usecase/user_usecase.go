package usecase

import (
	"errors"
	"simple-CRUD/pkg/entity"
	"simple-CRUD/pkg/repository"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

func CreateUser(user *entity.User) error {
	existingUser := repository.GetUserByEmail(user.Email)
	if existingUser != nil {
		return ErrEmailAlreadyExists
	}
	repository.CreateUser(user)
	return nil
}

func GetAllUsers() []entity.User {
	return repository.GetAllUsers()
}

func GetUserByID(id int) *entity.User {
	return repository.GetUserByID(id)
}

func UpdateUser(id int, user *entity.User) error {
	existingUser := repository.GetUserByEmail(user.Email)
	if existingUser != nil {
		return ErrEmailAlreadyExists
	}
	if isUpdated := repository.UpdateUser(id, user); !isUpdated {
		return errors.New("user not found")
	}
	return nil
}

func DeleteUser(id int) error {
	return repository.DeleteUser(id)
}
