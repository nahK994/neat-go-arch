package main

import (
	"fmt"
	"simple-CRUD/pkg/app"
	"simple-CRUD/pkg/handler"
	"simple-CRUD/pkg/repository"
	"simple-CRUD/pkg/router"
	"simple-CRUD/pkg/usecase"
)

func main() {
	config := app.GetConfig()

	db, _ := repository.MigrateDB(&config.DB)
	usecase := usecase.NewUserUsecase(db)
	handler := handler.NewUserHandler(usecase)

	r := router.SetupRouter(handler)

	addr := fmt.Sprintf("%s:%d", config.REST.Domain, config.REST.Port)
	r.Run(addr)
}
