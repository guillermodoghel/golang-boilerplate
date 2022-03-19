package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"guillermodoghel/golang-boilerplate/internal/logger"
)

type TimeController struct {
	logger *logrus.Logger
}

func NewTimeController() *TimeController {
	return &TimeController{
		logger: logger.GetLogger(),
	}
}

func (pc *TimeController) Time(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"time": time.Now().Format(time.RFC3339),
	})
}
