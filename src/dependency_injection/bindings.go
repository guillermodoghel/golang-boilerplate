package dependency_injection

import (
	"gorm.io/gorm"

	"guillermodoghel/golang-boilerplate/src/controllers"
	"guillermodoghel/golang-boilerplate/src/server"
	"guillermodoghel/golang-boilerplate/src/services/ping"
)

func Inject(db *gorm.DB) server.PingController {
	return controllers.NewPingController(ping.NewPingService(db))
}
