package handler

import (
	"net/http"
	"strconv"

	"simple-CRUD/pkg/entity"
	"simple-CRUD/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user *entity.User
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.ValidateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, usecase.GetAllUsers())
}

func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := usecase.GetUserByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user *entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.ValidateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.UpdateUser(id, user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := usecase.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
