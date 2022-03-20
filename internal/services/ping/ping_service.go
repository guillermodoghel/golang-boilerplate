package ping

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"guillermodoghel/golang-boilerplate/internal/logger"
	"guillermodoghel/golang-boilerplate/internal/services"
)

type PingService struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewPingService(db *gorm.DB) *PingService {
	return &PingService{
		db:     db,
		logger: logger.GetLogger(),
	}
}

func (s *PingService) Ping() (*services.PingResponse, error) {
	var result string
	err := s.db.Raw(`select "pong"`).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return &services.PingResponse{
		Message: result,
	}, nil
}
