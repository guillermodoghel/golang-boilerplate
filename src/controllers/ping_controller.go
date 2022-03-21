package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"guillermodoghel/golang-boilerplate/src/rest"
	"guillermodoghel/golang-boilerplate/src/server"
	"guillermodoghel/golang-boilerplate/src/services"
)

type PingController struct {
	pingService services.PingService
	logger      *logrus.Logger
}

func NewPingController(pingService services.PingService) *PingController {
	return &PingController{
		pingService: pingService,
		logger:      server.GetLogger(),
	}
}

func (pc *PingController) Ping(c *gin.Context) {
	ping, err := pc.pingService.Ping()
	if err != nil {
		pc.logger.Error("[PingController] ", err)
		rest.NewStatusInternalServerError(c)
		return
	}
	c.JSON(http.StatusOK, ping)
}
