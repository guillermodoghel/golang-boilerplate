package main

import (
	"github.com/jmoiron/sqlx"

	"guillermodoghel/golang-boilerplate/internal/controllers"
	"guillermodoghel/golang-boilerplate/internal/server"
	"guillermodoghel/golang-boilerplate/internal/services/ping"
)

func doRouterBindings(db *sqlx.DB) (server.PingController, server.TimeController) {
	return controllers.NewPingController(ping.NewPingService(db)),
		controllers.NewTimeController()
}
