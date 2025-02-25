package router

import (
	"simple-CRUD/pkg/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users", handler.CreateUser)
	r.GET("/users", handler.GetAllUsers)
	r.GET("/users/:id", handler.GetUserByID)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	return r
}
