package entity

import (
	"errors"
	"regexp"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	IsAdmin  bool   `json:"is_admin"`
	Password string `json:"password":omitempty`
}

type UserListResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

type UserResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	IsAdmin bool   `json:"is_admin"`
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func ValidateUser(user *User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Age < 18 {
		return errors.New("age must be greater than 18")
	}
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}
	return nil
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenerateTokenRequest struct {
	Id      int
	IsAdmin bool
}

type LoginInfo struct {
	Id       int
	Password string
	IsAdmin  bool
}
