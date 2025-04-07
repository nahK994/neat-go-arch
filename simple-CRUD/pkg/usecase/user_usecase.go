package usecase

import (
	"simple-CRUD/pkg/entity"
	"simple-CRUD/pkg/errors"
)

type UserRepository interface {
	CreateUser(name, email, hashedPassword string, age int, isAdmin bool) (entity.UserId, error)
	GetAllUsers() ([]entity.UserListResponse, error)
	GetUserByID(id int) (*entity.UserResponse, error)
	UpdateUser(id int, name, email string, age int) error
	DeleteUser(id int) error
	EmailExists(email string) (bool, error)
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

func (u *UserUsecase) CreateUser(user *entity.User) (int, error) {
	if err := entity.ValidateUser(user); err != nil {
		return -1, err
	}

	existingUser, _ := u.repo.EmailExists(user.Email)
	if existingUser {
		return -1, errors.ErrEmailAlreadyExists
	}

	id, err := u.repo.CreateUser(user.Name, user.Email, user.Password, user.Age, user.IsAdmin)
	if err != nil {
		return -1, err
	}

	return int(id), nil
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

	existingUser, err := u.repo.EmailExists(user.Email)
	if existingUser {
		return errors.ErrEmailAlreadyExists
	}
	if err != nil {
		return err
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
		return nil, errors.ErrEmailNotExists
	}

	if !checkPasswordHash(payload.Password, loginInfo.Password) {
		return nil, errors.ErrUnauthorized
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
