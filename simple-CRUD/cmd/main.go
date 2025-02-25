package main

import (
	"simple-CRUD/pkg/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
