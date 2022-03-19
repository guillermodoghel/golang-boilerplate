package ping

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"guillermodoghel/golang-boilerplate/internal/logger"
	"guillermodoghel/golang-boilerplate/internal/services"
)

type PingService struct {
	db *sqlx.DB
	logger *logrus.Logger
}

func NewPingService(db *sqlx.DB) *PingService {
	return &PingService{
		db: db,
		logger:logger.GetLogger(),
	}
}

func (s *PingService) Ping() (*services.PingResponse, error) {
	var result string
	err := s.db.QueryRow(`select "pong"`).Scan(&result)
	if err != nil {
		return nil, err
	}

	return &services.PingResponse{
		Message: result,
	}, nil
}
