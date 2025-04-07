package handler

import (
	"net/http"
	"strconv"

	"simple-CRUD/pkg/entity"
	"simple-CRUD/pkg/usecase"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

type Userhandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *Userhandler {
	return &Userhandler{
		usecase: usecase,
	}
}

func (h *Userhandler) Login(c *gin.Context) {
	var payload entity.LoginRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	loginInfo, err := h.usecase.Login(&payload)
	if err != nil {
		if err.Error() == usecase.ErrEmailAlreadyExists.Error() {
			c.JSON(http.StatusBadRequest, "email already exists")
			return
		} else if err.Error() == usecase.ErrUnauthorized.Error() {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		} else {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
	}

	c.JSON(http.StatusOK, loginInfo)
}

func (h *Userhandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password, _ = hashPassword(user.Password)
	if err := h.usecase.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := entity.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Age:     user.Age,
		IsAdmin: user.IsAdmin,
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *Userhandler) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, h.usecase.GetAllUsers())
}

func (h *Userhandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Userhandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user *entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.UpdateUser(id, user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "updated")
}

func (h *Userhandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
