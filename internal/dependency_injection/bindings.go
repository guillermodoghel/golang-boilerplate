package dependency_injection

import (
	"gorm.io/gorm"

	"guillermodoghel/golang-boilerplate/internal/controllers"
	"guillermodoghel/golang-boilerplate/internal/server"
	"guillermodoghel/golang-boilerplate/internal/services/ping"
)

func Inject(db *gorm.DB) (server.PingController, server.TimeController) {
	return controllers.NewPingController(ping.NewPingService(db)),
		controllers.NewTimeController()
}
