package server

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
		logger.SetOutput(os.Stdout)
		logger.SetLevel(logrus.DebugLevel)
	}
	return logger
}
