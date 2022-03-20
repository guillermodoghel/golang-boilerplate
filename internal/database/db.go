package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	mysqlUrlFmt  = "%s:%s@(%s:%s)/%s?parseTime=true&charset=utf8mb4"
	maxIdleConns = 20
	maxOpenConns = 150
)

func GetDB() (*gorm.DB, error) {
	var (
		dbUser     = os.Getenv("RDS_USER")
		dbPassword = os.Getenv("RDS_PASSWORD")
		dbHost     = os.Getenv("RDS_HOSTNAME")
		dbPort     = os.Getenv("RDS_PORT")
		dbName     = os.Getenv("RDS_DATABASE")
	)
	databaseUrl := fmt.Sprintf(mysqlUrlFmt, dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{
		Logger: dbLogger(),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	return db, nil
}

func dbLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  dbLoggerLevel(),
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
}

func dbLoggerLevel() logger.LogLevel {
	if os.Getenv("GIN_MODE") == "release" {
		return logger.Error
	} else {
		return logger.Info
	}
}
