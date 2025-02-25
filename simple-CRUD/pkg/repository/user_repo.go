package repository

import (
	"errors"
	"simple-CRUD/pkg/entity"
	"sync"
)

var (
	users  = []entity.User{}
	nextID = 1
	mu     sync.Mutex
)

func CreateUser(user *entity.User) {
	mu.Lock()
	defer mu.Unlock()
	user.ID = nextID
	nextID++
	users = append(users, *user)
}

func GetAllUsers() []entity.User {
	mu.Lock()
	defer mu.Unlock()
	return users
}

func GetUserByID(id int) *entity.User {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func GetUserByEmail(email string) *entity.User {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}

func UpdateUser(id int, updatedUser *entity.User) bool {
	mu.Lock()
	defer mu.Unlock()
	for i, user := range users {
		if user.ID == id {
			users[i] = *updatedUser
			users[i].ID = id
			return true
		}
	}
	return false
}

func DeleteUser(id int) error {
	mu.Lock()
	defer mu.Unlock()
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return errors.New("user not found")
		}
	}
	return nil
}
