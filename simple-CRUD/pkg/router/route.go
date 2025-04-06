package router

import (
	"simple-CRUD/pkg/handler"
	"simple-CRUD/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handler.Userhandler) *gin.Engine {
	r := gin.Default()

	r.POST("/users", handler.CreateUser)
	r.POST("/login", handler.Login)

	auth := r.Group("/users")
	auth.Use(middleware.AuthMiddleware(), middleware.AuthorizeUserMiddleware())
	{
		auth.GET("/:id", handler.GetUserByID)
		auth.PUT("/:id", handler.UpdateUser)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/users", handler.GetAllUsers)
		admin.DELETE("/users/:id", handler.DeleteUser)
	}

	return r
}
