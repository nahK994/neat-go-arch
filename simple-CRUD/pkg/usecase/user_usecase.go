package usecase

import (
	"errors"
	"simple-CRUD/pkg/entity"
)

var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrUnauthorized = errors.New("unauthorized")

type Password string
type IsAdmin bool

type UserRepository interface {
	CreateUser(name, email, hashedPassword string, age int, isAdmin bool) error
	GetAllUsers() ([]entity.UserListResponse, error)
	GetUserByID(id int) (*entity.UserResponse, error)
	UpdateUser(id int, name, email string, age int) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (*entity.User, error)
	GetLoginInfoByEmail(email string) (*entity.LoginInfo, error)
}

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) CreateUser(user *entity.User) error {
	if err := entity.ValidateUser(user); err != nil {
		return err
	}

	existingUser, _ := u.repo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return ErrEmailAlreadyExists
	}
	return u.repo.CreateUser(user.Name, user.Email, user.Password, user.Age, user.IsAdmin)
}

func (u *UserUsecase) GetAllUsers() []entity.UserListResponse {
	users, _ := u.repo.GetAllUsers()
	return users
}

func (u *UserUsecase) GetUserByID(id int) (*entity.UserResponse, error) {
	user, err := u.repo.GetUserByID(id)
	return user, err
}

func (u *UserUsecase) UpdateUser(id int, user *entity.User) error {
	if err := entity.ValidateUser(user); err != nil {
		return err
	}

	existingUser, err := u.repo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return ErrEmailAlreadyExists
	}
	if err != nil {
		return ErrEmailAlreadyExists
	}
	if err := u.repo.UpdateUser(id, user.Name, user.Email, user.Age); err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}

func (u *UserUsecase) Login(payload *entity.LoginRequest) (*entity.LoginResponse, error) {
	loginInfo, err := u.repo.GetLoginInfoByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if !checkPasswordHash(payload.Password, loginInfo.Password) {
		return nil, ErrUnauthorized
	}

	accessToken, err1 := generateJWT(loginInfo.IsAdmin, loginInfo.Id)
	if err1 != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		Id:          loginInfo.Id,
		IsAdmin:     loginInfo.IsAdmin,
		AccessToken: accessToken,
	}, nil
}
