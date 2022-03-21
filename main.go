package main

import (
	"flag"

	"guillermodoghel/golang-boilerplate/src/database"
	"guillermodoghel/golang-boilerplate/src/dependency_injection"
	"guillermodoghel/golang-boilerplate/src/server"
)

func main() {
	logger := server.GetLogger()
	logger.Info("🚀 Starting server")

	db, err := database.GetDB()
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	skipMigration := flag.Bool("skip-migration", false, "do not run migration scripts on start")
	flag.Parse()
	if !(*skipMigration) {
		if err := database.ExecuteMigrations(db); err != nil {
			panic(err)
		}
	}

	router := server.NewRouter()
	router.SetupRoutes(dependency_injection.Inject(db))

	router.Run()
}
