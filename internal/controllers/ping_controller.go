package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"guillermodoghel/golang-boilerplate/internal/logger"
	"guillermodoghel/golang-boilerplate/internal/rest"
	"guillermodoghel/golang-boilerplate/internal/services"
)

type PingController struct {
	pingService services.PingService
	logger      *logrus.Logger
}

func NewPingController(pingService services.PingService) *PingController {
	return &PingController{
		pingService: pingService,
		logger:      logger.GetLogger(),
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
