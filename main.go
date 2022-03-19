package main

import (
	"guillermodoghel/golang-boilerplate/internal/database"
	"guillermodoghel/golang-boilerplate/internal/dependency_injection"
	"guillermodoghel/golang-boilerplate/internal/logger"
	"guillermodoghel/golang-boilerplate/internal/server"
)

func main() {
	logger := logger.GetLogger()
	logger.Info("🚀 Starting server")

	db, err := database.GetDB()
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	router := server.NewRouter()
	router.SetupRoutes(dependency_injection.Inject(db))

	router.Run()
}
