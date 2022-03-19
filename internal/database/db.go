package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	mysqlUrlFmt  = "%s:%s@(%s:%s)/%s?parseTime=true&charset=utf8mb4"
	maxIdleConns = 20
	maxOpenConns = 150
)

func GetDB() (*sqlx.DB, error) {
	var (
		dbUser     = os.Getenv("RDS_USER")
		dbPassword = os.Getenv("RDS_PASSWORD")
		dbHost     = os.Getenv("RDS_HOSTNAME")
		dbPort     = os.Getenv("RDS_PORT")
		dbName     = os.Getenv("RDS_DATABASE")
	)
	databaseUrl := fmt.Sprintf(mysqlUrlFmt, dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect("mysql", databaseUrl)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	return db, nil
}
