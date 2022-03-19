package main

import (
	"guillermodoghel/golang-boilerplate/internal/database"
	"guillermodoghel/golang-boilerplate/internal/server"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}

	router := server.NewRouter()
	router.SetupRoutes(doRouterBindings(db))

	router.Run()
}
